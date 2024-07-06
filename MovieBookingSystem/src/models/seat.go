package models

import "fmt"

type SeatCategory string

const (
	SILVER   SeatCategory = "SILVER"
	GOLD     SeatCategory = "GOLD"
	PLATINUM SeatCategory = "PLATINUM"
)

type Seat struct {
	id           string
	row          string
	number       int
	seatCategory SeatCategory
}

func NewSeat(row string, number int, seatCategory SeatCategory) *Seat {
	return &Seat{row: row, number: number, seatCategory: seatCategory}
}

func (s *Seat) GetID() string {
	return s.id
}

func (s *Seat) GetRow() string {
	return s.row
}

func (s *Seat) GetNumber() int {
	return s.number
}

func (s *Seat) String() string {
	return fmt.Sprintf("row: %s, number: %d, seatCategory: %s", s.row, s.number, s.seatCategory)
}
