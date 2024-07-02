package main

import "fmt"

type Channel struct {
	name        string
	observerMap map[string]Observer
}

func NewChannel(name string) *Channel {
	return &Channel{name: name}
}

func (c *Channel) subscribe(observer Observer) {
	if c.observerMap == nil {
		c.observerMap = make(map[string]Observer)
	}
	c.observerMap[observer.getID()] = observer
}

func (c *Channel) unsubscribe(observer Observer) {
	delete(c.observerMap, observer.getID())
}

func (c *Channel) notifyAll(videoName string) {
	for _, observer := range c.observerMap {
		observer.update(c.name, videoName)
	}
}

func (c *Channel) uploadVideo(videoName string) {
	fmt.Printf("New video uploaded: %s on channel: %s\n", videoName, c.name)
	c.notifyAll(videoName)
}
