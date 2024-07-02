package main

import "fmt"

type Board struct {
	size  int
	board [][]IPiece
}

func NewBoard(size int) *Board {
	board := &Board{size: size}

	board.board = make([][]IPiece, size)

	for i := range board.board {
		board.board[i] = make([]IPiece, size)
	}

	return board
}

func (board *Board) getSize() int {
	return board.size
}

func (board *Board) getBoard() [][]IPiece {
	return board.board
}

func (board *Board) AddPieceToPosition(row int, col int, piece IPiece) bool {
	if row >= board.size || col >= board.size {
		return false
	}

	if board.board[row][col] != nil {
		return false
	}

	board.board[row][col] = piece

	return true
}

func (board *Board) isFull() bool {

	for row := range board.board {
		for col := range board.board[row] {
			if board.board[row][col] == nil {
				return false
			}
		}
	}

	return true
}
func (board *Board) print() {
	for row := range board.board {
		for col := range board.board[row] {
			printSymbol := " "
			if board.board[row][col] != nil {
				printSymbol = board.board[row][col].getPieceType().String()
			}
			fmt.Printf("| %s |", printSymbol)
		}
		fmt.Println()
	}
}
