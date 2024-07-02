package main

func main() {

	x := NewPieceX()
	o := NewPieceO()

	player1 := NewPlayer("Ravi", x)
	player2 := NewPlayer("Shruti", o)
	player3 := NewPlayer("Player", o)

	game := NewGame(2, 3)

	game.addPlayer(player1)
	game.addPlayer(player2)
	game.addPlayer(player3)

	game.start()
}
