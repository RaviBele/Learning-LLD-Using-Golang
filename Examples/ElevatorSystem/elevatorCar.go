package main

import "fmt"

type State int

const (
	MOVING State = iota + 1
	IDLE
)

func (s State) String() string {
	return [...]string{"MOVING", "IDLE"}[s-1]
}

type Elevator struct {
	id               int
	display          *Display
	currentFloor     int
	currentDirection Direction
	state            State
	button           *InternalButton
	door             *Door
}

func NewElevator(id int) *Elevator {
	return &Elevator{
		id:               id,
		display:          NewDisplay(),
		currentFloor:     0,
		currentDirection: NONE,
		state:            IDLE,
		button:           NewInternalButton(),
		door:             NewDoor(),
	}
}

func (elevator *Elevator) getID() int {
	return elevator.id
}

func (elevator *Elevator) move(floor int, direction Direction) {
	elevator.currentDirection = direction
	elevator.currentFloor = floor
	fmt.Printf("Elevator %d is moving to %s\n", elevator.id, direction)
	fmt.Printf("Elevator %d Floor: %d\n", elevator.id, floor)

	elevator.door.open()
	elevator.door.close()

	elevator.display.setDisplayMessage(elevator.currentFloor, direction)
	elevator.display.show()
}

func (elevator *Elevator) pressButton(floor int) {
	direction := NONE
	if floor > elevator.currentFloor {
		direction = UP
	} else if floor < elevator.currentFloor {
		direction = DOWN
	}
	elevator.button.press(floor, direction, elevator.id)
}

func (elevator *Elevator) getCurrentFloor() int {
	return elevator.currentFloor
}

func (elevator *Elevator) getCurrentDirection() Direction {
	return elevator.currentDirection
}

func (elevator *Elevator) Button() *InternalButton {
	return elevator.button
}
