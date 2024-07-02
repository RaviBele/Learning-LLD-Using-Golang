package main

import "fmt"

func main() {
	adidasFactory := GetSportsFactory("adidas")
	nikeFactory := GetSportsFactory("nike")

	adidasShoes := adidasFactory.makeShoes()
	nikeShoes := nikeFactory.makeShoes()

	adidasShirts := adidasFactory.makeShirts()
	nikeShirts := nikeFactory.makeShirts()

	printShoeDetails(adidasShoes)
	printShoeDetails(nikeShoes)
	printShirtDetails(adidasShirts)
	printShirtDetails(nikeShirts)
}

func printShoeDetails(s IShoesFactory) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirtsFactory) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
