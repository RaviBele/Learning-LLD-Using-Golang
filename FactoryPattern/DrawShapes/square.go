package main

import "fmt"

type Square struct {
	name string
}

func NewSquare(name string) *Square {
	return &Square{name: name}
}

func (t *Square) draw() {
	fmt.Printf("Drawn Square\n")
}

func (t *Square) getName() string {
	return t.name
}
