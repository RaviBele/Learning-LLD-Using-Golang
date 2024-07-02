package main

import "fmt"

type Rectangle struct {
	name string
}

func NewRectangle(name string) *Rectangle {
	return &Rectangle{name: name}
}

func (t *Rectangle) draw() {
	fmt.Printf("Drawn Rectangle\n")
}

func (t *Rectangle) getName() string {
	return t.name
}
