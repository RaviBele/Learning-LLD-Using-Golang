package main

type ElectricCarVehicle struct {
	Vehicle
}

func NewElectricCarVehicle(vehicleNumber string, withDisabled bool) IVehicle {
	return &ElectricCarVehicle{Vehicle: Vehicle{
		vehicleType:   ElectricCar,
		vehicleNumber: vehicleNumber,
		withDisabled:  withDisabled,
	}}
}