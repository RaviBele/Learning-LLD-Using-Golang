package main

import "github.com/google/uuid"

type MotorCycleParkingSpot struct {
	ParkingSpot
}

func NewMotorCycleParkingSpot() IParkingSpot {
	return &MotorCycleParkingSpot{
		ParkingSpot: ParkingSpot{
			parkingSpotID:   uuid.NewString(),
			isSpotAvailable: true,
			parkingSpotType: MotorCycle,
		},
	}
}
