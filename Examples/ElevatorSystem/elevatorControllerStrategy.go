package main

type ElevatorControllerStrategy interface {
	moveElevator(elevatorController *ElevatorController)
	addPendingRequest(pendingRequest *PendingRequest)
	SetElevatorControllers(controllers []*ElevatorController)
}
