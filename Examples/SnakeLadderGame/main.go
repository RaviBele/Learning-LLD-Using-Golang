package main

func main() {
	board := NewBoard(10, 10)
	board.AddSnakes(5)
	board.AddLadders(5)
	game := NewGame(board, 1)

	game.addPlayer("Ravi")
	game.addPlayer("Shruti")

	game.startGame()
}

