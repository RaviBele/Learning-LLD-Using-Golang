package main

import "fmt"

type ElevatorController struct {
	elevator *Elevator
}

func NewElevatorController(elevator *Elevator) *ElevatorController {
	return &ElevatorController{
		elevator: elevator,
	}
}

func (c *ElevatorController) getElevatorID() int {
	return c.elevator.getID()
}

func (c *ElevatorController) getElevator() *Elevator {
	return c.elevator
}

func (c *ElevatorController) acceptNewRequest(floorID int, dir Direction) {
	GetElevatorSystem().GetElevatorControllerStrategy().addPendingRequest(NewPendingRequest(floorID, dir))
	c.controlElevator()
}

func (c *ElevatorController) controlElevator() {
	fmt.Println("Moving elevator....")
	GetElevatorSystem().GetElevatorControllerStrategy().moveElevator(c)
}

func (c *ElevatorController) moveElevator(floor int, direction Direction) {
	c.elevator.move(floor, direction)
}

func (c *ElevatorController) getCurrentFloor() int {
	return c.elevator.getCurrentFloor()
}

func (c *ElevatorController) getCurrentDirection() Direction {
	return c.elevator.getCurrentDirection()
}
