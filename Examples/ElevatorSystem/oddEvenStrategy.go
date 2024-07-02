package main

import (
	"fmt"
	"math/rand"
)

type OddEvenStragey struct {
	elevatorControllers []*ElevatorController
}

func NewOddEvenStrategy() *OddEvenStragey {
	return &OddEvenStragey{}
}

func (e *OddEvenStragey) SetElevatorControllers(controllers []*ElevatorController) {
	e.elevatorControllers = controllers
}

func (e *OddEvenStragey) selectElevator(floor int, direction Direction) int {

	for _, controller := range e.elevatorControllers {
		if controller.getElevatorID()%2 == floor%2 {
			currentFloor := controller.getCurrentFloor()
			currentDirection := controller.getCurrentDirection()

			if floor < currentFloor && currentDirection == UP {
				return controller.getElevatorID()
			} else if floor > currentFloor && currentDirection == DOWN {
				return controller.getElevatorID()
			} else {
				fmt.Printf("accepting NewRequest elevator....%d %s\n", currentFloor, currentDirection)
				return controller.getElevatorID()
			}
		}
	}

	return rand.Intn(len(e.elevatorControllers))
}
