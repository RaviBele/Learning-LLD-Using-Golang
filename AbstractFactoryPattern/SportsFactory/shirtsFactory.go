package main

type IShirtsFactory interface {
	setLogo(string)
	getLogo() string
	setSize(int)
	getSize() int
}
