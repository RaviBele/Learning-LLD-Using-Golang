package main

import "github.com/google/uuid"

type DisabledParkingSpot struct {
	ParkingSpot
}

func NewDisabledParkingSpot() IParkingSpot {
	return &DisabledParkingSpot{
		ParkingSpot: ParkingSpot{
			parkingSpotID:   uuid.NewString(),
			isSpotAvailable: true,
			parkingSpotType: Disabled,
		},
	}
}
