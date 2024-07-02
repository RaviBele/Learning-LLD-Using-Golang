package main

type PendingRequest struct {
	floor     int
	direction Direction
}

func NewPendingRequest(floor int, direction Direction) *PendingRequest {
	return &PendingRequest{floor: floor, direction: direction}
}

func (r *PendingRequest) getFloor() int { return r.floor }

func (r *PendingRequest) getDirection() Direction { return r.direction }
