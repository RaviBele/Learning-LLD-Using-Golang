package models

import (
	"time"

	"github.com/google/uuid"
)

type Show struct {
	id            string
	startTime     time.Time
	movie         *Movie
	screen        *Screen
	bookedSeatIDs map[string]bool
}

func NewShow(startTime time.Time, movie *Movie, screen *Screen) *Show {
	return &Show{
		id:            uuid.NewString(),
		startTime:     startTime,
		movie:         movie,
		screen:        screen,
		bookedSeatIDs: make(map[string]bool),
	}
}

func (s *Show) GetStartTime() time.Time {
	return s.startTime
}

func (s *Show) GetMovie() *Movie {
	return s.movie
}

func (s *Show) GetBookedSeatIDs() map[string]bool {
	return s.bookedSeatIDs
}

func (s *Show) GetAvailableSeats() []*Seat {
	seats := s.screen.GetAllSeats()
	availableSeats := make([]*Seat, 0, len(seats))
	for _, seat := range seats {
		if _, ok := s.bookedSeatIDs[seat.GetID()]; !ok {
			availableSeats = append(availableSeats, seat)
		}
	}
	return availableSeats
}

func (s *Show) UpdateBookedSeats(seats []*Seat) {
	for _, seat := range seats {
		s.bookedSeatIDs[seat.GetID()] = true
	}
}
