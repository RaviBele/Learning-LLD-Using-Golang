package main

type IShoesFactory interface {
	setLogo(string)
	getLogo() string
	setSize(int)
	getSize() int
}
