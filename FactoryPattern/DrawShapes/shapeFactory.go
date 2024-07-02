package main

import "fmt"

type Shape interface {
	draw()
	getName() string
}

func getShape(name string) (Shape, error) {
	if name == "triangle" {
		return NewTriangle(name), nil
	}

	if name == "rectangle" {
		return NewRectangle(name), nil
	}

	if name == "square" {
		return NewSquare(name), nil
	}

	return nil, fmt.Errorf("Wrong shape is passed")
}