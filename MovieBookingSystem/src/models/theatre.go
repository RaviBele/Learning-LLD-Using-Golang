package models

import "github.com/google/uuid"

type Theatre struct {
	id      string
	name    string
	city    string
	screens []*Screen
	shows   []*Show
}

func NewTheatre(name string, city string) *Theatre {
	return &Theatre{
		id:   uuid.NewString(),
		name: name,
		city: city,
	}
}

func (a *Theatre) AddScreen(screen *Screen) {
	a.screens = append(a.screens, screen)
}

func (a *Theatre) RemoveScreen(screen *Screen) {
	for i, v := range a.screens {
		if v == screen {
			a.screens = append(a.screens[:i], a.screens[i+1:]...)
		}
	}
}

func (a *Theatre) AddShow(show *Show) {
	a.shows = append(a.shows, show)
}

func (a *Theatre) RemoveShow(show *Show) {
	for i, v := range a.shows {
		if v == show {
			a.shows = append(a.shows[:i], a.shows[i+1:]...)
		}
	}
}

func (a *Theatre) GetShows() []*Show {
	return a.shows
}

func (a *Theatre) GetScreens() []*Screen {
	return a.screens
}

func (a *Theatre) GetShowsByMovieName(movieName string) []*Show {
	shows := []*Show{}
	for _, v := range a.shows {
		if v.movie.GetName() == movieName {
			shows = append(shows, v)
		}
	}

	return shows
}

func (a *Theatre) GetAvailableSeatsForShow(show *Show) []*Seat {
	return show.GetAvailableSeats()
}

func (a *Theatre) MakeBooking(show *Show, seats []*Seat) *Booking {
	booking := NewBooking(show, seats, NewPayment())

	show.UpdateBookedSeats(seats)

	return booking
}
