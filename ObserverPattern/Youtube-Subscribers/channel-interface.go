package main

type ChannelInterface interface {
	subscribe(Observer)
	unsubscribe(Observer)
	uploadVideo(string)
	notifyAll(string)
}
