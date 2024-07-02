package main

import "github.com/google/uuid"

type LargeParkingSpot struct {
	ParkingSpot
}

func NewLargeParkingSpot() IParkingSpot {
	return &LargeParkingSpot{
		ParkingSpot: ParkingSpot{
			parkingSpotID:   uuid.NewString(),
			isSpotAvailable: true,
			parkingSpotType: Large,
		},
	}
}
