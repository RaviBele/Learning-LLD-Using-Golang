package main

import "github.com/google/uuid"

type Player struct {
	id              string
	name            string
	currentPosition int
}

func NewPlayer(name string) *Player {
	return &Player{
		id:              uuid.NewString(),
		name:            name,
		currentPosition: -1,
	}
}

func (p *Player) getName() string {
	return p.name
}

func (p *Player) getCurrentPosition() int {
	return p.currentPosition
}

func (p *Player) updateCurrentPosition(position int) {
	p.currentPosition = position
}
