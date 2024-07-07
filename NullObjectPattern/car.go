package main

type Car struct {
	Vehicle
}

func NewCar(tankCapacity int, seatingCapacity int) *Car {
	return &Car{
		Vehicle: Vehicle{
			tankCapacity:    tankCapacity,
			seatingCapacity: seatingCapacity,
		},
	}
}

func (c *Car) GetTankCapacity() int {
	return c.tankCapacity
}

func (c *Car) GetSeatingCapacity() int {
	return c.seatingCapacity
}
