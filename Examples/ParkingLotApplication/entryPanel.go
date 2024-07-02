package main

import (
	"fmt"

	"github.com/google/uuid"
)

type EntryPanel struct {
	entryPanelID string
}

func NewEntryPanel() *EntryPanel {
	return &EntryPanel{
		entryPanelID: uuid.NewString(),
	}
}

func (p *EntryPanel) getID() string {
	return p.entryPanelID
}

func (p *EntryPanel) getParkingTicket(vehicle IVehicle) (*ParkingTicket, error) {
	if !isValidVehicleType(vehicle.getVehicleType()) {
		fmt.Println("Parking is not allowed")
		return nil, fmt.Errorf("Parking is not allowed")
	}

	parkingLot := getParkingLot()

	var availableSpot IParkingSpot
	var availableFloor *ParkingFloor
	for _, floor := range parkingLot.getListOfParkingFloors() {
		spot, err := floor.getAvailableParkingSpot(vehicle)
		if err != nil {
			fmt.Printf("Parking not available on floor %s\n", floor.getParkingFloorID())
		} else {
			availableSpot = spot
			availableFloor = floor
		}
	}

	if availableSpot == nil {
		return nil, fmt.Errorf("Parking is not available")
	}

	ticket := p.generateParkingTicket(vehicle, availableFloor.getParkingFloorID(), availableSpot.getParkingSpotID(), availableSpot.getParkingSpotType())

	availableSpot.assignSpotToVehicle(&vehicle)

	return ticket, nil
}

func (p *EntryPanel) generateParkingTicket(vehicle IVehicle, floorID string, spotID string, spotType ParkingSpotType) *ParkingTicket {
	return NewParkingTicket(vehicle.getVehicleType(), vehicle.getVehicleNumber(), spotID, floorID, spotType)
}
