package main

import "math/rand"

type Dice struct {
	diceCount int
}

func NewDice(diceCount int) *Dice {
	return &Dice{diceCount: diceCount}
}

func (d *Dice) rollRice() int {
	diceValue := 0
	minDiceValue := 1
	maxDiceValue := 6
	for i := 0; i < d.diceCount; i++ {
		diceValue += rand.Intn(maxDiceValue-minDiceValue+1) + minDiceValue
	}

	return diceValue
}
