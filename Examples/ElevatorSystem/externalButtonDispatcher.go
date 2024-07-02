package main

import "fmt"

type ExternalButtonDispatcher struct {
	elevatorControllers []*ElevatorController
}

func NewExternalButtonDispatcher() *ExternalButtonDispatcher {
	return &ExternalButtonDispatcher{}
}

func (externalButtonDispatcher *ExternalButtonDispatcher) addElevatorController(controller *ElevatorController) {
	externalButtonDispatcher.elevatorControllers = append(externalButtonDispatcher.elevatorControllers, controller)
}

func (externalButtonDispatcher *ExternalButtonDispatcher) submitRequest(floor int, direction Direction) {
	elevatorID := GetElevatorSystem().GetElevatorSelectionStrategy().selectElevator(floor, direction)
	fmt.Println("Elevator ID: ", elevatorID)
	for _, controller := range externalButtonDispatcher.elevatorControllers {
		if controller.getElevatorID() == elevatorID {
			controller.acceptNewRequest(floor, direction)
		}
	}
}
