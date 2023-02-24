package constants

const (
	SPEC_TYPE_STRING = "STRING"
	SPEC_TYPE_INT    = "INT"
)

type ProductSpecification string

const (
	TYRE_SPEC_EAN          ProductSpecification = "EAN_CODE"
	TYRE_SPEC_NAME         ProductSpecification = "NAME"
	TYRE_SPEC_REFERENCE    ProductSpecification = "REFERENCE"
	TYRE_SPEC_WIDTH        ProductSpecification = "WIDTH"
	TYRE_SPEC_ASPECT_RATIO ProductSpecification = "ASPECT_RATIO"
	TYRE_SPEC_CONSTRUCTION ProductSpecification = "CONSTRUCTION"
	TYRE_SPEC_RIM          ProductSpecification = "RIM"
	TYRE_SPEC_LOAD         ProductSpecification = "LOAD"
	TYRE_SPEC_SPEED        ProductSpecification = "SPEED"
	TYRE_SPEC_SEASON       ProductSpecification = "SEASON"
	TYRE_SPEC_EPREL_ID     ProductSpecification = "EPREL_ID"
)
