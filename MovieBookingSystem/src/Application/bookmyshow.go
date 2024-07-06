package application

import (
	"moviebookingsystem/controller"
	"moviebookingsystem/models"
	"time"
)

type BookMyShow struct {
	movieController   *controller.MovieController
	theatreController *controller.TheatreController
}

func BookMyShowApplication() *BookMyShow {
	return &BookMyShow{
		movieController:   controller.NewMovieController(),
		theatreController: controller.NewTheaterController(),
	}
}

func (b *BookMyShow) AddMovie(city string, movieName string, duration float64) *models.Movie {
	movie := models.NewMovie(movieName, duration)
	b.movieController.AddMovieToCity(city, movie)
	return movie
}

func (b *BookMyShow) GetAllMoviesByCity(city string) []*models.Movie {
	return b.movieController.GetMoviesByCity(city)
}

func (b *BookMyShow) GetAllMovies() []*models.Movie {
	return b.movieController.GetAllMovies()
}

func (b *BookMyShow) AddTheatreToCity(city string, name string) *models.Theatre {
	theatre := models.NewTheatre(name, city)
	b.theatreController.AddTheatreToCity(city, theatre)
	return theatre
}

func (b *BookMyShow) GetAllTheatersByCity(city string) []*models.Theatre {
	return b.theatreController.GetTheatersByCity(city)
}

func (b *BookMyShow) GetAllTheaters() []*models.Theatre {
	return b.theatreController.GetAllTheaters()
}

func (b *BookMyShow) AddScreen(theatre *models.Theatre, screenSize int, goldSeat int, silverSeats int, platinumSeats int, number int) *models.Screen {
	screen := models.NewScreen(screenSize, goldSeat, silverSeats, platinumSeats, number)
	b.theatreController.AddScreenToTheater(theatre, screen)
	return screen
}

func (b *BookMyShow) AddShow(theatre *models.Theatre, startTime time.Time, screen *models.Screen, movie *models.Movie) *models.Show {
	show := models.NewShow(startTime, movie, screen)
	b.theatreController.AddShowToTheater(theatre, show)
	return show
}

func (b *BookMyShow) GetAllShowsByMovieName(theatre *models.Theatre, movieName string) []*models.Show {
	return b.theatreController.GetAllShowsByMovieName(theatre, movieName)
}

func (b *BookMyShow) GetAvailableSeatsForShow(theatre *models.Theatre, show *models.Show) []*models.Seat {
	return b.theatreController.GetAvailableSeatsForShow(theatre, show)
}

func (b *BookMyShow) MakeBooking(theatre *models.Theatre, show *models.Show, seats []*models.Seat) *models.Booking {
	return b.theatreController.MakeBooking(theatre, show, seats)
}
