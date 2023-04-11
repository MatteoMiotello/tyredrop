package converters

import (
	"pillowww/titw/graph/model"
	"pillowww/titw/models"
)

func VehicleTypeToGraphQL(vehicleType *models.VehicleType) *model.VehicleType {
	name := vehicleType.R.VehicleTypeLanguages[0].Name

	return &model.VehicleType{
		ID:   vehicleType.ID,
		Code: vehicleType.Code,
		Name: name,
	}
}
