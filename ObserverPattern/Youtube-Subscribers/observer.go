package main


type Observer interface {
	update(string, string)
	getID() string
}