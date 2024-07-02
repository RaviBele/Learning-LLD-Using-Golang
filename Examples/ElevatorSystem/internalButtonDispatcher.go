package main

type internalButtonDispatcher struct {
	elevatorControllers []*ElevatorController
}

func NewInternalButtonDispatcher() *internalButtonDispatcher {
	return &internalButtonDispatcher{}
}

func (e *internalButtonDispatcher) addElevatorController(controller *ElevatorController) {
	e.elevatorControllers = append(e.elevatorControllers, controller)
}

func (e *internalButtonDispatcher) submitRequest(floorID int, direction Direction, elevatorID int) {
	for _, controller := range e.elevatorControllers {
		if controller.getElevatorID() == elevatorID {
			controller.acceptNewRequest(floorID, direction)
		}
	}
}
