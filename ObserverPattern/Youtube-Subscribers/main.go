package main

func main() {
	
	ytChannel := NewChannel("CodingChannel")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	ytChannel.subscribe(observerFirst)
	ytChannel.subscribe(observerSecond)

	ytChannel.uploadVideo("Golang")
}
