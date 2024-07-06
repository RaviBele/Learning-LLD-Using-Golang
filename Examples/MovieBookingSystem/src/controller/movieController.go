package controller

import "moviebookingsystem/models"

type MovieController struct {
	cityToMovieMapping map[string]map[*models.Movie]bool
	movies             map[*models.Movie]bool
}

func NewMovieController() *MovieController {
	return &MovieController{
		cityToMovieMapping: make(map[string]map[*models.Movie]bool),
		movies:             make(map[*models.Movie]bool),
	}
}

func (c *MovieController) AddMovieToCity(city string, movie *models.Movie) {
	if _, ok := c.cityToMovieMapping[city]; !ok {
		c.cityToMovieMapping[city] = make(map[*models.Movie]bool)
	}
	c.cityToMovieMapping[city][movie] = true
	c.movies[movie] = true
}

func (c *MovieController) GetMoviesByCity(city string) []*models.Movie {
	movies := make([]*models.Movie, len(c.cityToMovieMapping[city]))

	for k, _ := range c.cityToMovieMapping[city] {
		movies = append(movies, k)
	}

	return movies
}

func (c *MovieController) GetAllMovies() []*models.Movie {
	movies := make([]*models.Movie, len(c.movies))

	for k, _ := range c.movies {
		movies = append(movies, k)
	}

	return movies
}
