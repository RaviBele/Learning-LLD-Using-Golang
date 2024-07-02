package main

type Subject interface {
	register(observer Observer)
	unregisterObserver(observer Observer)
	updateAvailability()
}