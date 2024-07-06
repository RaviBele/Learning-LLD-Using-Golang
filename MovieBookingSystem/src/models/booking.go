package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Booking struct {
	id          string
	show        *Show
	bookedSeats []*Seat
	payment     *Payment
}

func NewBooking(show *Show, bookedSeats []*Seat, payment *Payment) *Booking {
	return &Booking{
		id:          uuid.NewString(),
		show:        show,
		bookedSeats: bookedSeats,
		payment:     payment,
	}
}

func (b *Booking) String() string {
	bookingDetails := fmt.Sprintf("Booking ID: %s\nMovieName: %s\nScreen: %d\nSeats: \n",
		b.id, b.show.GetMovie().GetName(), b.show.screen.number)

	for _, s := range b.bookedSeats {
		bookingDetails += fmt.Sprintf("%v\n", s)
	}

	return bookingDetails
}
