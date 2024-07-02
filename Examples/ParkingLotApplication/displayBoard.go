package main

import "fmt"

type DisplayBoard struct {}

func NewDisplayBoard() *DisplayBoard {
	return &DisplayBoard{}
}

func (d *DisplayBoard) displayMessage(spotType ParkingSpotType, freeSpots int) {
	fmt.Printf("%s spots free: %d\n", spotType, freeSpots)
}