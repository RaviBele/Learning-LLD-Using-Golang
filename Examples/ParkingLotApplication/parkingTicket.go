package main

import (
	"time"

	"github.com/google/uuid"
)

type ParkingTicket struct {
	parkingTicketID string
	vehicleType     VehicleType
	vehicleNumber   string
	parkingSpotID   string
	parkingFloorID  string
	parkingSpotType ParkingSpotType
	startTime       time.Time
	endTime         time.Time
	amount          int
}

func NewParkingTicket(vehicleType VehicleType, vehicleNumber string, parkingSpotID string, parkingFloorID string, spotType ParkingSpotType) *ParkingTicket {
	return &ParkingTicket{
		parkingTicketID: uuid.NewString(),
		vehicleType:     vehicleType,
		vehicleNumber:   vehicleNumber,
		parkingSpotID:   parkingSpotID,
		parkingFloorID:  parkingFloorID,
		parkingSpotType: spotType,
	}
}

func (t *ParkingTicket) setStartTime() {
	t.startTime = time.Now()
}

func (t *ParkingTicket) setEndTime() {
	t.endTime = time.Now()
}

func (t *ParkingTicket) setAmount(amount int) {
	t.amount = amount
}

func (t *ParkingTicket) getStartTime() time.Time {
	return t.startTime
}

func (t *ParkingTicket) getEndTime() time.Time {
	return t.endTime
}

func (t *ParkingTicket) getAmount() int {
	return t.amount
}

func (t *ParkingTicket) getVehicleType() VehicleType {
	return t.vehicleType
}

func (t *ParkingTicket) getParkingSpotID() string {
	return t.parkingSpotID
}

func (t *ParkingTicket) getParkingFloorID() string {
	return t.parkingFloorID
}

func (t *ParkingTicket) getVehicleNumber() string {
	return t.vehicleNumber
}
