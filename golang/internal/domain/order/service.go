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

func (s *service) createOrderRowFromCart(ctx context.Context, currency *models.Currency, order *models.Order, cart *models.Cart) error {
	p, err := s.itemPriceDao.
		Load(models.ProductItemPriceRels.ProductItem).
		FindOneById(ctx, cart.ProductItemPriceID)

	if err != nil {
		return err
	}

	productItem := p.R.ProductItem

	if p.R.ProductItem.SupplierQuantity < cart.Quantity {
		return errors.New("Quantity not available")
	}

	row := &models.OrderRow{
		OrderID:            order.ID,
		ProductItemPriceID: p.ID,
		Quantity:           cart.Quantity,
		Amount:             p.Price * cart.Quantity,
	}

	err = s.orderDao.Insert(ctx, row)

	if err != nil {
		return err
	}

	productItem.SupplierQuantity = productItem.SupplierQuantity - cart.Quantity

	err = s.itemDao.Update(ctx, productItem)

	if err != nil {
		return err
	}

	return nil
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

	for _, cart := range carts {
		err = s.createOrderRowFromCart(ctx, currentCurrency, newOrder, cart)

		if err != nil {
			return nil, err
		}
	}

	return newOrder, nil
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
