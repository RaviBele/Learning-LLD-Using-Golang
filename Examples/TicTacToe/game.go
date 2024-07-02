package main

import "fmt"

type Game struct {
	numOfPlayers int
	players      []*Player
	board        *Board
}

func NewGame(numOfPlayers int, boardSize int) *Game {
	game := &Game{numOfPlayers: numOfPlayers, board: NewBoard(boardSize)}
	return game
}

func (game *Game) addPlayer(player *Player) {
	if len(game.players) == game.numOfPlayers {
		fmt.Printf("%d players added, No more players required...\n", game.numOfPlayers)
		return
	}

	for i := range game.players {
		if game.players[i] == player {
			return
		}
	}
	game.players = append(game.players, player)
}

func (game *Game) start() {
	if game.numOfPlayers < len(game.players) {
		fmt.Printf("Number of players required: %d, players added: %d\n", game.numOfPlayers, len(game.players))
		return
	}

	noWinner := true

	fmt.Println("Starting game")
	for noWinner {
		game.board.print()
		currentPlayer := game.players[0]

		if game.board.isFull() {
			noWinner = false
			continue
		}

		fmt.Printf("%s: Select your position: ", currentPlayer.getName())
		fmt.Println()
		var row, col int
		_, err := fmt.Scan(&row, &col)
		if err != nil {
			fmt.Println("Selected wrong position... try again.")
			continue
		}
		isSuccessFullyAdded := game.board.AddPieceToPosition(row, col, currentPlayer.getPlayingPiece())

		if !isSuccessFullyAdded {
			fmt.Println("Selected wrong position... try again.")
			continue
		}

		if game.hasWinner(row, col, currentPlayer.getPlayingPiece()) {
			game.board.print()
			fmt.Printf("%s is winner!\n", currentPlayer.getName())
			return
		}

		game.players = append(game.players[1:], currentPlayer)
	}

	fmt.Println("Game is tie")
}

func (game *Game) hasWinner(row, col int, playingPiece IPiece) bool {

	rowMatched := true
	colMatched := true
	diagonalMatched := true
	reverseDiagonal := true

	for i := 0; i < game.board.getSize(); i++ {
		if game.board.getBoard()[row][i] == nil || game.board.getBoard()[row][i] != playingPiece {
			rowMatched = false
		}
	}

	for i := 0; i < game.board.getSize(); i++ {
		if game.board.getBoard()[i][col] == nil || game.board.getBoard()[i][col] != playingPiece {
			colMatched = false
		}
	}

	j := 0
	for i := 0; i < game.board.getSize(); i++ {
		if game.board.getBoard()[i][j] == nil || game.board.getBoard()[i][j] != playingPiece {
			diagonalMatched = false
		}
		j++
	}

	j = len(game.board.getBoard()[0]) - 1
	for i := 0; i < game.board.getSize(); i++ {
		if game.board.getBoard()[i][j] == nil || game.board.getBoard()[i][j] != playingPiece {
			reverseDiagonal = false
		}
		j--
	}

	return rowMatched || colMatched || diagonalMatched || reverseDiagonal

}
