package main

import "github.com/google/uuid"

type ElectricCarParkingSpot struct {
	ParkingSpot
}

func NewElectricCarParkingSpot() IParkingSpot {
	return &ElectricCarParkingSpot{
		ParkingSpot: ParkingSpot{
			parkingSpotID:   uuid.NewString(),
			isSpotAvailable: true,
			parkingSpotType: ElectricCarSpot,
		},
	}
}
