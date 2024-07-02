package main

import "github.com/google/uuid"

type CarParkingSpot struct {
	ParkingSpot
}

func NewCarParkingSpot() IParkingSpot {
	return &CarParkingSpot{
		ParkingSpot: ParkingSpot{
			parkingSpotID:   uuid.NewString(),
			isSpotAvailable: true,
			parkingSpotType: Compact,
		},
	}
}
