package main

import (
	"fmt"
	"math"

	"github.com/google/uuid"
)

type ExitPanel struct {
	exitPanelID string
}

func NewExitPanel() *ExitPanel {
	return &ExitPanel{
		exitPanelID: uuid.NewString(),
	}
}

func (p *ExitPanel) getID() string {
	return p.exitPanelID
}

func (p *ExitPanel) checkout(ticket *ParkingTicket) (*ParkingTicket, error) {
	totalDurationInHours := p.calculateDurationInHours(ticket)
	vacatedSpot := getParkingLot().vacateParkingSpot(ticket.parkingSpotID, ticket.parkingFloorID, ticket.parkingSpotType)

	if vacatedSpot == nil {
		return ticket, fmt.Errorf("Unable to locate parking spot for ticket %v", ticket)
	}
	totalAmount := p.calculatePrice(ticket.parkingSpotType, totalDurationInHours)

	ticket.setAmount(totalAmount)
	return ticket, nil
}

func (p *ExitPanel) calculatePrice(spotType ParkingSpotType, totalHours int) int {
	hourlyRate, _ := getHourlyCostObject().getCost(spotType)
	if totalHours == 0 {
		return hourlyRate
	}
	fmt.Printf("Total Hours: %v and hourlyRate: %v\n", totalHours, hourlyRate)
	return totalHours * hourlyRate
}

func (p *ExitPanel) calculateDurationInHours(ticket *ParkingTicket) int {
	ticket.setEndTime()
	startTime := ticket.getStartTime()
	endTime := ticket.getEndTime()

	duration := endTime.Sub(startTime)

	hours := duration.Hours()
	return int(math.Floor(hours))
}
