package main

import "fmt"

type Game struct {
	board   *Board
	players []*Player
	dice    *Dice
}

func NewGame(board *Board, diceCount int) *Game {
	return &Game{board: board, dice: NewDice(diceCount)}
}

func (g *Game) addPlayer(name string) {
	g.players = append(g.players, NewPlayer(name))
}

func (g *Game) getPlayers() []*Player {
	return g.players
}

func (g *Game) startGame() {
	if len(g.players) < 2 {
		fmt.Printf("Atleast two players are required...\n")
		return
	}

	rank := 1
	for len(g.players) > 1 {
		player := g.players[0]
		g.players = g.players[1:]
		fmt.Printf("Player %v from positon: %d Rolling dice...\n", player.getName(), player.getCurrentPosition())
		diceValue := g.dice.rollRice()
		fmt.Printf("Player %v: Dice value: %d\n", player.getName(), diceValue)

		if player.getCurrentPosition()+diceValue > g.board.getEnd() {
			fmt.Printf("Dice value is out of board end. Skipping dice roll\n")
			g.players = append(g.players, player)
		} else if player.getCurrentPosition()+diceValue == g.board.getEnd() {
			fmt.Printf("Player %v: reached end. Ranked: %d\n", player.getName(), rank)
			rank += 1
		} else {
			player.updateCurrentPosition(g.board.movePeiceAndGetFinalPosition(player.getCurrentPosition() + diceValue))
			fmt.Printf("Player %v: moved to position: %d\n", player.getName(), player.getCurrentPosition())
			g.players = append(g.players, player)
		}
	}

	fmt.Printf("Player %v lost...\n", g.players[0].getName())
}
