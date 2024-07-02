package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	cells [][]*Cell
	end   int
	n     int
	m     int
}

func NewBoard(n int, m int) *Board {
	board := &Board{end: (n * m) - 1, n: n, m: m}

	board.cells = make([][]*Cell, n)
	for i := 0; i < n; i++ {
		board.cells[i] = make([]*Cell, m)
		for j := 0; j < m; j++ {
			board.cells[i][j] = NewCell(nil)
		}
	}

	return board
}

func (board *Board) getEnd() int {
	return board.end
}

func (board *Board) AddSnakes(snakeCount int) {
	minSnakeTailValue := 2
	maxSnakeMouthValue := board.end - 2
	i := 0
	for i < snakeCount {
		snakeMouth := rand.Intn(maxSnakeMouthValue-(minSnakeTailValue+1)+1) + (minSnakeTailValue + 1)
		snakeTail := rand.Intn(snakeMouth-minSnakeTailValue+1) + minSnakeTailValue
		snakeHeadCellRow := int(snakeMouth / board.n)
		snakeHeadCellCol := (snakeMouth % board.m)
		if !board.cells[snakeHeadCellRow][snakeHeadCellCol].hasJump() {
			board.cells[snakeHeadCellRow][snakeHeadCellCol].AssignJump(NewJump(snakeMouth, snakeTail))
			i++
		}
		fmt.Printf("Added snake: %d\n", i)
	}
}

func (board *Board) AddLadders(ladderCount int) {
	minLadderStartValue := 2
	maxLadderEndValue := board.end - 2
	i := 0
	for i < ladderCount {
		ladderEnd := rand.Intn(maxLadderEndValue-(minLadderStartValue+1)+1) + (minLadderStartValue + 1)
		ladderStart := rand.Intn(ladderEnd-minLadderStartValue+1) + minLadderStartValue
		ladderStartCellRow := int(ladderStart / board.n)
		ladderStartCellCol := (ladderStart % board.m)
		if !board.cells[ladderStartCellRow][ladderStartCellCol].hasJump() {
			board.cells[ladderStartCellRow][ladderStartCellCol].AssignJump(NewJump(ladderStart, ladderEnd))
			i++
		}
		fmt.Printf("Added ladder: %d\n", i)
	}
}

func (b *Board) movePeiceAndGetFinalPosition(desiredPosition int) int {
	cellRow := int(desiredPosition / b.n)
	cellCol := (desiredPosition % b.m)

	if b.cells[cellRow][cellCol].hasJump() {
		if b.cells[cellRow][cellCol].getJump().getEnd() > b.cells[cellRow][cellCol].getJump().getStart() {
			fmt.Printf("%d is ladder starting position %d, \nclimbing ladder till ladder end position: %d\n", desiredPosition, b.cells[cellRow][cellCol].getJump().getStart(), b.cells[cellRow][cellCol].getJump().getEnd())
		} else {
			fmt.Printf("%d is snake mouth position at %d, \nsnake ate you and moved to end position: %d\n", desiredPosition, b.cells[cellRow][cellCol].getJump().getStart(), b.cells[cellRow][cellCol].getJump().getEnd())
		}
		return b.cells[cellRow][cellCol].getJump().getEnd()
	}

	return desiredPosition
}
