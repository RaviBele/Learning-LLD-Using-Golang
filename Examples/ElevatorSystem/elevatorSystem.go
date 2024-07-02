package main

type ElevatorSystem struct {
	elevatorControllers       []*ElevatorController
	elevatorControlStrategy   ElevatorControllerStrategy
	elevatorSelectionStrategy ElevatorSelectionStrategy
	floors                    []*Floor
}

var elevatorSystem *ElevatorSystem

func GetElevatorSystem() *ElevatorSystem {
	if elevatorSystem != nil {
		return elevatorSystem
	}
	elevatorSystem = &ElevatorSystem{}
	return elevatorSystem
}

func (elevatorSystem *ElevatorSystem) addFloor(floor *Floor) {
	elevatorSystem.floors = append(elevatorSystem.floors, floor)
}

func (elevatorSystem *ElevatorSystem) addElevator(elevator *ElevatorController) {
	elevatorSystem.elevatorControllers = append(elevatorSystem.elevatorControllers, elevator)
	elevatorSystem.elevatorSelectionStrategy.SetElevatorControllers(elevatorSystem.elevatorControllers)
	elevatorSystem.elevatorControlStrategy.SetElevatorControllers(elevatorSystem.elevatorControllers)
	for _, floor := range elevatorSystem.floors {
		floor.Button().addController(elevator)
	}

	for _, controller := range elevatorSystem.elevatorControllers {
		controller.getElevator().Button().addController(elevator)
	}
}

func (elevatorSystem *ElevatorSystem) removeElevator(elevator *ElevatorController) {
	for i, elevatorController := range elevatorSystem.elevatorControllers {
		if elevatorController.getElevatorID() == elevator.getElevatorID() {
			elevatorSystem.elevatorControllers = append(elevatorSystem.elevatorControllers[:i], elevatorSystem.elevatorControllers[i+1:]...)
		}
	}

	elevatorSystem.elevatorSelectionStrategy.SetElevatorControllers(elevatorSystem.elevatorControllers)
	elevatorSystem.elevatorControlStrategy.SetElevatorControllers(elevatorSystem.elevatorControllers)
}

func (elevatorSystem *ElevatorSystem) setElevatorSelectionStrategy(elevatorSelectionStrategy ElevatorSelectionStrategy) {
	elevatorSystem.elevatorSelectionStrategy = elevatorSelectionStrategy
}

func (elevatorSystem *ElevatorSystem) setElevatorControllerStrategy(elevatorControllerStrategy ElevatorControllerStrategy) {
	elevatorSystem.elevatorControlStrategy = elevatorControllerStrategy
}

func (elevatorSystem *ElevatorSystem) GetElevatorSelectionStrategy() ElevatorSelectionStrategy {
	return elevatorSystem.elevatorSelectionStrategy
}

func (elevatorSystem *ElevatorSystem) GetElevatorControllerStrategy() ElevatorControllerStrategy {
	return elevatorSystem.elevatorControlStrategy
}

func (elevatorSystem *ElevatorSystem) getControllers() []*ElevatorController {
	return elevatorSystem.elevatorControllers
}

func (elevatorSystem *ElevatorSystem) getFloors() []*Floor {
	return elevatorSystem.floors
}
