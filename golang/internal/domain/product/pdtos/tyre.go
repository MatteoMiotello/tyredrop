package pdtos

import (
	"fmt"
	"pillowww/titw/pkg/constants"
	"strconv"
)

type Tyre struct {
	EANCode     string
	ProductName string
	Reference   string
	Brand       string
	Season      string
	Price       string
	EprelID     string
	Quantity    int
	VehicleType constants.VehicleType
	TyreDimension
}

type TyreDimension struct {
	Width        int
	AspectRatio  int
	Construction string
	Rim          int
	Load         int
	Speed        string
}

func (t *Tyre) GetProductName() string {
	return t.Reference
}

func (t *Tyre) GetVehicleType() constants.VehicleType {
	return t.VehicleType
}

func (t *Tyre) GetProductCode() string {
	return t.EANCode
}

func (t *Tyre) GetSupplierProductPrice() string {
	return t.Price
}

func (t *Tyre) GetProductCategoryCode() constants.ProductCategoryType {
	return constants.PRODUCT_CATEGORY_TYRE
}

func (t *Tyre) GetSpecifications() map[constants.ProductSpecification]string {
	return map[constants.ProductSpecification]string{
		constants.TYRE_SPEC_EAN:          t.EANCode,
		constants.TYRE_SPEC_NAME:         t.ProductName,
		constants.TYRE_SPEC_REFERENCE:    t.Reference,
		constants.TYRE_SPEC_WIDTH:        strconv.Itoa(t.Width),
		constants.TYRE_SPEC_ASPECT_RATIO: strconv.Itoa(t.AspectRatio),
		constants.TYRE_SPEC_CONSTRUCTION: t.Construction,
		constants.TYRE_SPEC_RIM:          strconv.Itoa(t.Rim),
		constants.TYRE_SPEC_LOAD:         strconv.Itoa(t.Load),
		constants.TYRE_SPEC_SPEED:        t.Speed,
		constants.TYRE_SPEC_SEASON:       t.Season,
		constants.TYRE_SPEC_EPREL_ID:     t.EprelID,
	}
}

func (t *Tyre) GetBrandName() string {
	return t.Brand
}

func (t *Tyre) Validate() bool {
	if t.Width == 0 {
		return false
	}

	if t.Load == 0 {
		return false
	}

	if t.Construction == "" {
		return false
	}

	if t.Brand == "" {
		return false
	}

	if t.EANCode == "" {
		return false
	}

	return true
}

func (t *Tyre) GetSupplierProductQuantity() int {
	return t.Quantity
}

func (t *Tyre) BuildName() string {
	return fmt.Sprintf("%d/%d %s %d %d %s", t.Width, t.AspectRatio, t.Construction, t.Rim, t.Load, t.Speed)
}

func (t *Tyre) GetEprelProductCode() *string {
	fmt.Println(t.EprelID)

	if len(t.EprelID) == 0 {
		return nil
	}

	i := t.EprelID
	return &i
}
