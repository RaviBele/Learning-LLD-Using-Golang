package models

import "github.com/google/uuid"

type Screen struct {
	id     string
	number int
	size   int
	seats  []*Seat
}

func NewScreen(size int, noOfGoldSeats int, noOfSilverSeats int, noOfPlatinumSeats int, number int) *Screen {
	screen := &Screen{
		id:     uuid.NewString(),
		size:   size,
		number: number,
	}

	ch := 'a'
	noOfSeatsPerRow := 10
	row := 1
	for i := 1; i <= noOfPlatinumSeats; i++ {
		screen.seats = append(screen.seats, NewSeat(string(ch), row, PLATINUM))
		if row == noOfSeatsPerRow {
			row = 1
			ch += 1
		}
		row += 1
	}

	row = 1
	for i := 1; i <= noOfGoldSeats; i++ {
		screen.seats = append(screen.seats, NewSeat(string(ch), row, GOLD))
		if row == noOfSeatsPerRow {
			row = 1
			ch++
		}
		row += 1
	}

	row = 1
	for i := 1; i <= noOfSilverSeats; i++ {
		screen.seats = append(screen.seats, NewSeat(string(ch), row, SILVER))
		if row == noOfSeatsPerRow {
			row = 1
			ch++
		}
		row += 1
	}

	return screen
}

func (s *Screen) GetSize() int {
	return s.size
}

func (s *Screen) GetAllSeats() []*Seat {
	return s.seats
}
