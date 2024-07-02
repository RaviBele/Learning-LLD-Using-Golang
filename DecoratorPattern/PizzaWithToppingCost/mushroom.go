package main

import "fmt"

type Mushroom struct {
	pizza BasePizza
}

func NewMushroom(basePizza BasePizza) *Mushroom {
	fmt.Printf("Added Mushrooms\n")
	return &Mushroom{pizza: basePizza}
}

func (e *Mushroom) cost() int {
	return e.pizza.cost() + 20
}

func (e *Mushroom) getName() string {
	return e.pizza.getName() + " with Mushroom"
}
