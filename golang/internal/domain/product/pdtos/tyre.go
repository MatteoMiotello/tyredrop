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
	RunFlat     bool
	VehicleType constants.VehicleType
	ImageUrl    *string
	TyreDimension
	EprelSpec
}

type TyreDimension struct {
	Width        int
	AspectRatio  int
	Construction string
	Rim          float64
	Load         int
	Speed        string
}

type EprelSpec struct {
	FuelEfficiency            string
	WetGripClass              string
	ExternalRollingNoiseClass string
	ExternalRollingNoiseLevel int
	LoadVersion               string
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
	rollingNoise := ""
	if t.ExternalRollingNoiseLevel != 0 {
		rollingNoise = strconv.Itoa(t.ExternalRollingNoiseLevel)
	}

	return map[constants.ProductSpecification]string{
		constants.TYRE_SPEC_EAN:                          t.EANCode,
		constants.TYRE_SPEC_NAME:                         t.ProductName,
		constants.TYRE_SPEC_REFERENCE:                    t.Reference,
		constants.TYRE_SPEC_WIDTH:                        strconv.Itoa(t.Width),
		constants.TYRE_SPEC_ASPECT_RATIO:                 strconv.Itoa(t.AspectRatio),
		constants.TYRE_SPEC_CONSTRUCTION:                 t.Construction,
		constants.TYRE_SPEC_RIM:                          strconv.FormatFloat(t.Rim, 'f', 0, 64),
		constants.TYRE_RUNFLAT:                           strconv.FormatBool(t.RunFlat),
		constants.TYRE_SPEC_LOAD:                         strconv.Itoa(t.Load),
		constants.TYRE_SPEC_SPEED:                        t.Speed,
		constants.TYRE_SPEC_SEASON:                       t.Season,
		constants.TYRE_SPEC_EPREL_ID:                     t.EprelID,
		constants.TYRE_SPEC_FUEL_EFFICIENCY:              t.FuelEfficiency,
		constants.TYRE_SPEC_WET_GRIP_CLASS:               t.WetGripClass,
		constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_CLASS: t.ExternalRollingNoiseClass,
		constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_LEVEL: rollingNoise,
		constants.TYRE_SPEC_LOAD_VERSION:                 t.LoadVersion,
	}
}

func (t *Tyre) GetBrandName() string {
	return t.Brand
}

func (t *Tyre) Validate() bool {
	if t.Width < 1 {
		return false
	}

	if t.Rim < 1 {
		return false
	}

	if t.AspectRatio < 1 {
		return false
	}

	if t.Load < 1 {
		return false
	}

	if len(t.Construction) == 0 {
		return false
	}

	if len(t.Brand) == 0 {
		return false
	}

	if len(t.EANCode) == 0 {
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

func (t *Tyre) GetProductImageUrl() *string {
	return t.ImageUrl
}

func (t *Tyre) GetPriceAdditionCodes() []string {
	if t.GetVehicleType() == constants.VEHICLE_MOTO {
		return []string{
			constants.TYRE_PFU_MOTO.String(),
		}
	}

	if t.Rim == 17.5 {
		return []string{
			constants.TYRE_PFU_R_17_5.String(),
		}
	}

	if t.Rim <= 18 {
		return []string{
			constants.TYRE_PFU_R_13_18.String(),
		}
	}

	if t.Rim <= 20 {
		return []string{
			constants.TYRE_PFU_R_19_20.String(),
		}
	}

	if t.Rim <= 23 {
		return []string{
			constants.TYRE_PFU_R_21_23.String(),
		}
	}

	return []string{
		constants.TYRE_PFU_R_22_5_23_5.String(),
	}
}
