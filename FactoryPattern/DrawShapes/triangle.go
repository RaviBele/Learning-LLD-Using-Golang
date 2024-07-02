package main

import "fmt"

type Triangle struct {
	name string
}

func NewTriangle(name string) *Triangle {
	return &Triangle{name: name}
}

func (t *Triangle) draw() {
	fmt.Printf("Drawn triangle\n")
}

func (t *Triangle) getName() string {
	return t.name
}
