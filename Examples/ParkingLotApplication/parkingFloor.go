package main

import (
	"fmt"

	"github.com/google/uuid"
)

type ParkingFloor struct {
	parkingSpots   map[ParkingSpotType][]IParkingSpot
	parkingFloorID string
	displayBoard   *DisplayBoard
}

func NewParkingFloor() *ParkingFloor {
	return &ParkingFloor{
		parkingFloorID: uuid.NewString(),
		displayBoard:   NewDisplayBoard(),
		parkingSpots:   make(map[ParkingSpotType][]IParkingSpot),
	}
}

func (f *ParkingFloor) getParkingFloorID() string {
	return f.parkingFloorID
}

func (f *ParkingFloor) getListOfParkingSpots() map[ParkingSpotType][]IParkingSpot {
	return f.parkingSpots
}

func (f *ParkingFloor) getAvailableParkingSpot(vehicle IVehicle) (IParkingSpot, error) {

	vehicleSpotType := f.getSpotTypoForVehicle(vehicle.getVehicleType(), vehicle.isWithDisabled())
	for _, p := range f.parkingSpots[vehicleSpotType] {
		if p.isavailable() {
			return p, nil
		}
	}

	return nil, fmt.Errorf("Parking full for %s", vehicleSpotType)
}

func (f *ParkingFloor) getSpotTypoForVehicle(vehicleType VehicleType, isDisabled bool) ParkingSpotType {
	switch vehicleType {
	case Car:
		if isDisabled {
			return Disabled
		}
		return Compact
	case ElectricCar:
		if isDisabled {
			return Disabled
		}
		return ElectricCarSpot
	case MotorBike:
		if isDisabled {
			return Disabled
		}
		return MotorCycle
	default:
		if isDisabled {
			return Disabled
		}
		return Large
	}
}

func (f *ParkingFloor) showDisplayBoard() {
	for spotType, spots := range f.parkingSpots {
		freeSpots := 0
		for _, spot := range spots {
			if spot.isavailable() {
				freeSpots += 1
			}
		}
		f.displayBoard.displayMessage(spotType, freeSpots)
	}
}

func (f *ParkingFloor) inUseSpots(vehicleType VehicleType) []IParkingSpot {
	var parkingSpots []IParkingSpot
	for _, spot := range f.parkingSpots[f.getSpotTypoForVehicle(vehicleType, false)] {
		if !spot.isavailable() {
			parkingSpots = append(parkingSpots, spot)
		}
	}
	return parkingSpots
}

func (f *ParkingFloor) addParkingSpot(spot IParkingSpot) {
	spots := f.parkingSpots[spot.getParkingSpotType()]
	for _, spotItem := range spots {
		if spotItem.getParkingSpotID() == spot.getParkingSpotID() {
			return
		}
	}

	f.parkingSpots[spot.getParkingSpotType()] = append(f.parkingSpots[spot.getParkingSpotType()], spot)
}
