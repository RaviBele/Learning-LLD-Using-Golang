package main

type TruckVehicle struct {
	Vehicle
}

func NewTruckVehicle(vehicleNumber string, withDisabled bool) IVehicle {
	return &TruckVehicle{Vehicle: Vehicle{
		vehicleType:   Truck,
		vehicleNumber: vehicleNumber,
		withDisabled:  withDisabled,
	}}
}