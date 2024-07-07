package main

import "fmt"

func main() {

	vehicle := GetNewVehicle("Car")

	fmt.Printf("Car Tank capacity: %d\n", vehicle.GetTankCapacity())
	fmt.Printf("Car Seating capacity: %d\n", vehicle.GetSeatingCapacity())

	vehicle = GetNewVehicle("Bike")

	fmt.Printf("Bike Tank capacity: %d\n", vehicle.GetTankCapacity())
	fmt.Printf("Bike Seating capacity: %d\n", vehicle.GetSeatingCapacity())
}
