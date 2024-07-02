package main

import "fmt"

func main() {
	elevatorSystem := GetElevatorSystem()
	elevatorSystem.setElevatorSelectionStrategy(NewOddEvenStrategy())
	elevatorSystem.setElevatorControllerStrategy(NewLookAlgorithmStrategy())

	totalFloors := 50

	for i := 0; i < totalFloors; i++ {
		elevatorSystem.addFloor(NewFloor(i))
	}

	totalElevators := 4

	for i := 0; i < totalElevators; i++ {
		elevatorSystem.addElevator(NewElevatorController(NewElevator(i)))
	}

	//Request 1
	fmt.Println("Person at floor 1 presses UP Button")
	for _, floor := range elevatorSystem.getFloors() {
		if floor.getID() == 1 {
			floor.pressButton(UP)
		}
	}

	//Request 2
	fmt.Println("Person at floor 5 presses UP Button")
	for _, floor := range elevatorSystem.getFloors() {
		if floor.getID() == 5 {
			floor.pressButton(UP)
		}
	}

	//Request 3
	fmt.Println("Person presses 10 in elevator 2")
	for _, controller := range elevatorSystem.getControllers() {
		if controller.getElevatorID() == 2 {
			controller.getElevator().pressButton(10)
		}
	}

	//Request 4
	fmt.Println("Person presses 6 in elevator 2")
	for _, controller := range elevatorSystem.getControllers() {
		if controller.getElevatorID() == 2 {
			controller.getElevator().pressButton(6)
		}
	}

	//Request 5
	fmt.Println("Person at floor 7 presses DOWN Button")
	for _, floor := range elevatorSystem.getFloors() {
		if floor.getID() == 7 {
			floor.pressButton(DOWN)
		}
	}

	//Request 6
	fmt.Println("Person presses 1 in elevator 3")
	for _, controller := range elevatorSystem.getControllers() {
		if controller.getElevatorID() == 3 {
			controller.getElevator().pressButton(6)
		}
	}
}
