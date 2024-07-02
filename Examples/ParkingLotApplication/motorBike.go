package main

type MotorBikeVehicle struct {
	Vehicle
}

func NewMotorBikeVehicle(vehicleNumber string, withDisabled bool) IVehicle {
	return &MotorBikeVehicle{Vehicle: Vehicle{
		vehicleType:   MotorBike,
		vehicleNumber: vehicleNumber,
		withDisabled:  withDisabled,
	}}
}
