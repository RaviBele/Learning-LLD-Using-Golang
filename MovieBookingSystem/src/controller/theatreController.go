package controller

import "moviebookingsystem/models"

type TheatreController struct {
	cityToTheaterMapping map[string]map[*models.Theatre]struct{}
	theatres             map[*models.Theatre]struct{}
}

func NewTheaterController() *TheatreController {
	return &TheatreController{
		cityToTheaterMapping: make(map[string]map[*models.Theatre]struct{}),
		theatres:             make(map[*models.Theatre]struct{}),
	}
}

func (c *TheatreController) AddTheatreToCity(city string, theatre *models.Theatre) {
	if _, ok := c.cityToTheaterMapping[city]; !ok {
		c.cityToTheaterMapping[city] = make(map[*models.Theatre]struct{})
	}
	c.cityToTheaterMapping[city][theatre] = struct{}{}
	c.theatres[theatre] = struct{}{}
}

func (c *TheatreController) GetTheatersByCity(city string) []*models.Theatre {
	theatres := make([]*models.Theatre, len(c.cityToTheaterMapping[city]))

	for k, _ := range c.cityToTheaterMapping[city] {
		theatres = append(theatres, k)
	}

	return theatres
}

func (c *TheatreController) GetAllTheaters() []*models.Theatre {
	theatres := make([]*models.Theatre, len(c.theatres))

	for k, _ := range c.theatres {
		theatres = append(theatres, k)
	}

	return theatres
}

func (c *TheatreController) AddScreenToTheater(theater *models.Theatre, screen *models.Screen) {
	theater.AddScreen(screen)
}

func (c *TheatreController) AddShowToTheater(theater *models.Theatre, show *models.Show) {
	theater.AddShow(show)
}

func (c *TheatreController) GetAllShowsByMovieName(theater *models.Theatre, movieName string) []*models.Show {
	return theater.GetShowsByMovieName(movieName)
}

func (c *TheatreController) GetAvailableSeatsForShow(theater *models.Theatre, show *models.Show) []*models.Seat {
	return theater.GetAvailableSeatsForShow(show)
}

func (c *TheatreController) MakeBooking(theater *models.Theatre, show *models.Show, seats []*models.Seat) *models.Booking {
	return theater.MakeBooking(show, seats)
}
