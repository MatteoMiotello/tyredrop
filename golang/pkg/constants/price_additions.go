package constants

type PriceAddition string

const (
	TYRE_PFU_R_13_18     PriceAddition = "PFU_R_13_18"
	TYRE_PFU_R_19_20     PriceAddition = "PFU_R_19_20"
	TYRE_PFU_R_21_23     PriceAddition = "PFU_R_21_23"
	TYRE_PFU_R_17_5      PriceAddition = "PFU_R_17.5"
	TYRE_PFU_R_22_5_23_5 PriceAddition = "PFU_R_22.5_23.5"
	TYRE_PFU_MOTO        PriceAddition = "PFU_MOTO"
)

func (r PriceAddition) String() string {
	return string(r)
}
