package models

import "github.com/google/uuid"

type Movie struct {
	id       string
	name     string
	duration float64
}

func NewMovie(name string, duration float64) *Movie {
	return &Movie{
		id:       uuid.NewString(),
		name:     name,
		duration: duration,
	}
}

func (m *Movie) GetName() string {
	return m.name
}

func (m *Movie) GetDuration() float64 {
	return m.duration
}
