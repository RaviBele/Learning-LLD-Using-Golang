package main

type NullVehicle struct {
	Vehicle
}

func NewNullVehicle() *NullVehicle {
	return &NullVehicle{
		Vehicle: Vehicle{
			tankCapacity:    0,
			seatingCapacity: 0,
		},
	}
}

func (c *NullVehicle) GetTankCapacity() int {
	return c.tankCapacity
}

func (c *NullVehicle) GetSeatingCapacity() int {
	return c.seatingCapacity
}
