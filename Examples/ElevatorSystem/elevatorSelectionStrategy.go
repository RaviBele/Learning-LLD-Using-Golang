package main

type ElevatorSelectionStrategy interface {
	selectElevator(floor int, direction Direction) int
	SetElevatorControllers(controllers []*ElevatorController)
}
