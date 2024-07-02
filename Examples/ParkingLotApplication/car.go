package main

type CarVehicle struct {
	Vehicle
}

func NewCarVehicle(vehicleNumber string, withDisabled bool) IVehicle {
	return &CarVehicle{Vehicle: Vehicle{
		vehicleType:   Car,
		vehicleNumber: vehicleNumber,
		withDisabled:  withDisabled,
	}}
}