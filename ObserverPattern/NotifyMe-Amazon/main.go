package main

func main() {

	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.registerObserver(observerFirst)
	shirtItem.registerObserver(observerSecond)

	shirtItem.updateAvailability()
}
