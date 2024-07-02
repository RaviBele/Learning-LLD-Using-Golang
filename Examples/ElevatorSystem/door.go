package main

import "fmt"

type Door struct {
}

func NewDoor() *Door {
	return &Door{}
}

func (d *Door) open() {
	fmt.Println("Opening Door...")
}

func (d *Door) close() {
	fmt.Println("Closing Door...")
}
