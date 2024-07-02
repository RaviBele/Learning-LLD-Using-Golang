package main

import "sort"

type LookAlgorithmStrategy struct {
	pendingRequests     []*PendingRequest
	elevatorControllers []*ElevatorController
}

func NewLookAlgorithmStrategy() *LookAlgorithmStrategy {
	return &LookAlgorithmStrategy{}
}

func (s *LookAlgorithmStrategy) SetElevatorControllers(controllers []*ElevatorController) {
	s.elevatorControllers = controllers
}

func (s *LookAlgorithmStrategy) moveElevator(elevator *ElevatorController) {
	var currentRequest *PendingRequest

	up := []*PendingRequest{}
	down := []*PendingRequest{}

	currentDirection := elevator.getCurrentDirection()
	if currentDirection == NONE {
		if s.pendingRequests[0].getFloor() < elevator.getCurrentFloor() {
			currentDirection = DOWN
		} else {
			currentDirection = UP
		}
	}
	for i := 0; i < len(s.pendingRequests); i++ {
		if s.pendingRequests[i].getFloor() < elevator.getCurrentFloor() {
			down = append(down, s.pendingRequests[i])
		}

		if s.pendingRequests[i].getFloor() > elevator.getCurrentFloor() {
			up = append(up, s.pendingRequests[i])
		}
	}

	sort.Slice(down, func(i, j int) bool {
		return down[i].getFloor() > down[j].getFloor()
	})

	sort.Slice(up, func(i, j int) bool {
		return up[i].getFloor() < up[j].getFloor()
	})

	run := 2

	for i := 0; i < run; i++ {
		if currentDirection == DOWN {
			for j := 0; j < len(down); j++ {
				currentRequest = down[j]
				elevator.moveElevator(currentRequest.getFloor(), currentRequest.getDirection())
			}

			currentDirection = UP
		} else if currentDirection == UP {
			for j := 0; j < len(up); j++ {
				currentRequest = up[j]
				elevator.moveElevator(currentRequest.getFloor(), currentRequest.getDirection())
			}

			currentDirection = DOWN
		}
	}
}

func (s *LookAlgorithmStrategy) addPendingRequest(request *PendingRequest) {
	s.pendingRequests = append(s.pendingRequests, request)
}
