package main

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) update(channelName string, videoName string) {
	fmt.Printf("Sending email to customer: %s for new video upload on %s on channel %s\n", c.id, videoName, channelName)
}

func (c *Customer) getID() string {
	return c.id
}
