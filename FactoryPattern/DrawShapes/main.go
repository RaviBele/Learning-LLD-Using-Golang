package main

import "fmt"

func main() {
	triangle, _ := getShape("triangle")

	fmt.Printf("Shape: %v\n", triangle.getName())
	triangle.draw()

	rectangle, _ := getShape("rectangle")

	fmt.Printf("Shape: %v\n", rectangle.getName())
	rectangle.draw()

	square, _ := getShape("square")

	fmt.Printf("Shape: %v\n", square.getName())
	square.draw()
}
