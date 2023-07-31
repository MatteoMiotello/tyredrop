package order

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"pillowww/titw/graph/model"
	"pillowww/titw/internal/currency"
	"pillowww/titw/internal/domain/product"
	"pillowww/titw/models"
)

type service struct {
	orderDao     *Dao
	currencyDao  *currency.Dao
	itemDao      *product.ItemDao
	itemPriceDao *product.ItemPriceDao
}

func NewService(
	dao *Dao,
	currencyDao *currency.Dao,
	itemDao *product.ItemDao,
	itemPriceDao *product.ItemPriceDao,
) service {
	return service{
		dao,
		currencyDao,
		itemDao,
		itemPriceDao,
	}
}

func (s *service) createOrderRowFromCart(ctx context.Context, currency *models.Currency, order *models.Order, cart *models.Cart) (*models.OrderRow, error) {
	p, err := s.itemPriceDao.
		Load(models.ProductItemPriceRels.ProductItem).
		FindOneById(ctx, cart.ProductItemPriceID)

	if err != nil {
		return nil, err
	}

	productItem := p.R.ProductItem

	if p.R.ProductItem.SupplierQuantity < cart.Quantity {
		return nil, errors.New("Quantity not available")
	}

	additions, err := s.itemPriceDao.FindPriceAdditionsByProductItemPriceID(ctx, p.ID)

	if err != nil {
		return nil, err
	}

	priceWithAddition := 0

	for _, add := range additions {
		priceWithAddition = priceWithAddition + (add.AdditionValue * cart.Quantity)
	}

	row := &models.OrderRow{
		OrderID:            order.ID,
		ProductItemPriceID: p.ID,
		Quantity:           cart.Quantity,
		Amount:             p.Price * cart.Quantity,
		AdditionsAmount:    priceWithAddition,
	}

	err = s.orderDao.Insert(ctx, row)

	if err != nil {
		return nil, err
	}

	productItem.SupplierQuantity = productItem.SupplierQuantity - cart.Quantity

	err = s.itemDao.Update(ctx, productItem)

	if err != nil {
		return nil, err
	}

	return row, nil
}

func (s *service) CreateNewOrder(ctx context.Context, userBilling *models.UserBilling, address *models.UserAddress, carts models.CartSlice) (*models.Order, error) {
	currentCurrency, err := s.currencyDao.FindDefault(ctx)

	if err != nil {
		return nil, err
	}

	defaultTax, err := s.orderDao.FindDefaultTax(ctx)

	if err != nil {
		return nil, err
	}

	newOrder := &models.Order{
		CurrencyID:    currentCurrency.ID,
		TaxID:         defaultTax.ID,
		UserBillingID: userBilling.ID,
		AddressName:   address.AddressName,
		AddressLine1:  address.AddressLine1,
		AddressLine2:  address.AddressLine2,
		City:          address.City,
		PostalCode:    address.PostalCode,
		Province:      address.Province,
		Country:       address.Country,
		Status:        model.OrderStatusNew.String(),
	}

	err = s.orderDao.Insert(ctx, newOrder)

	if err != nil {
		return nil, err
	}

	var priceWithAdditions int
	var orderAmount int

	for _, cart := range carts {
		row, err := s.createOrderRowFromCart(ctx, currentCurrency, newOrder, cart)

		if err != nil {
			return nil, err
		}

		orderAmount = orderAmount + row.Amount
		priceWithAdditions = priceWithAdditions + (row.AdditionsAmount + row.Amount)
	}

	taxAmountFloat := (float64(defaultTax.MarkupPercentage) / 100) * float64(orderAmount)
	newOrder.PriceAmount = priceWithAdditions
	newOrder.TaxesAmount = int(taxAmountFloat)
	newOrder.PriceAmountTotal = newOrder.PriceAmount + newOrder.TaxesAmount

	err = s.orderDao.
		Update(ctx, newOrder)

	if err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (s *service) updateOrderPrice(ctx context.Context) {

}

func (s *service) updateOrderStatus(ctx context.Context, order *models.Order, newStatus model.OrderStatus) error {
	if !newStatus.IsValid() {
		return errors.New("Invalid status prompted")
	}

	if newStatus.String() == order.Status {
		return errors.New(fmt.Sprintf("The order already is in %s status", order.Status))
	}

	if !newStatus.IsValidForOrder(order) {
		return errors.New(fmt.Sprintf("%s can't be applied on current order with status: %s", newStatus.String(), order.Status))
	}

	order.Status = newStatus.String()

	err := s.orderDao.Insert(ctx, order)

	if err != nil {
		return err
	}

	return nil
}
