package main

import "fmt"

type ExtraCheese struct {
	pizza BasePizza
}

func NewExtraCheese(basePizza BasePizza) *ExtraCheese {
	fmt.Printf("Added Extra Cheese\n")
	return &ExtraCheese{pizza: basePizza}
}

func (e *ExtraCheese) cost() int {
	return e.pizza.cost() + 10
}

func (e *ExtraCheese) getName() string {
	return e.pizza.getName() + " with extra cheese"
}
