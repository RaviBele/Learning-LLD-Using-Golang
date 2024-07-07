package main

func GetNewVehicle(vehicleType string) IVehicle {
	switch vehicleType {
	case "Car":
		return NewCar(10, 15)
	}
	return NewNullVehicle()
}
