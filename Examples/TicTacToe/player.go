package main

type Player struct {
	name         string
	playingPiece IPiece
}

func NewPlayer(name string, playingPiece IPiece) *Player {
	return &Player{name: name, playingPiece: playingPiece}
}

func (p *Player) getName() string {
	return p.name
}

func (p *Player) getPlayingPiece() IPiece {
	return p.playingPiece
}
