package main

import "fmt"

type Direction int

const (
	UP Direction = iota + 1
	DOWN
	NONE
)

func (d Direction) String() string {
	return [...]string{"UP", "DOWN", "NONE"}[d-1]
}

type Display struct {
	currentFloor int
	direction    Direction
}

func NewDisplay() *Display {
	return &Display{
		direction: UP,
	}
}

func (d *Display) setDisplayMessage(currentFloor int, direction Direction) {
	d.currentFloor = currentFloor
	d.direction = direction
}

func (d *Display) show() {
	fmt.Printf("Going %s, current Floor: %d\n", d.direction, d.currentFloor)
}
