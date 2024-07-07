package main

type IVehicle interface {
	GetTankCapacity() int
	GetSeatingCapacity() int
}

type Vehicle struct {
	IVehicle
	tankCapacity    int
	seatingCapacity int
}
