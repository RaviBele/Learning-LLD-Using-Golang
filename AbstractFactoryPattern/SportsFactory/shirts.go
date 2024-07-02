package main

type Shirts struct {
	logo string
	size int
}

func (s *Shirts) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirts) getLogo() string {
	return s.logo
}

func (s *Shirts) setSize(size int) {
	s.size = size
}

func (s *Shirts) getSize() int {
	return s.size
}
