package constants

type VehicleType string

const (
	VEHICLE_CAR          VehicleType = "CAR"
	VEHICLE_MOTO         VehicleType = "MOTO"
	VEHICLE_TRUCK        VehicleType = "TRUCK"
	VEHICLE_QUAD         VehicleType = "QUAD"
	VEHICLE_AGRICULTURAL VehicleType = "AGRICULTURAL"
)

func (r VehicleType) String() string {
	return string(r)
}
