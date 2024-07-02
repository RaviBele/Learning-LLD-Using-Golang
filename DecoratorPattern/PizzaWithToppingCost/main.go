package main

import "fmt"

func main() {
	pizza := NewMargaritta()

	printPrice(pizza.getName(), pizza.cost())

	pizzaWithExtraCheese := NewExtraCheese(pizza)

	printPrice(pizzaWithExtraCheese.getName(), pizzaWithExtraCheese.cost())

	pizzaWithExtraCheeseAndMushroom := NewMushroom(pizzaWithExtraCheese)

	printPrice(pizzaWithExtraCheeseAndMushroom.getName(), pizzaWithExtraCheeseAndMushroom.cost())

	pizza1 := NewVegDelight()

	printPrice(pizza1.getName(), pizza1.cost())

	pizza1WithExtraCheese := NewExtraCheese(pizza1)

	printPrice(pizza1WithExtraCheese.getName(), pizza1WithExtraCheese.cost())

	pizza1WithExtraCheeseAndMushroom := NewMushroom(pizza1WithExtraCheese)

	printPrice(pizza1WithExtraCheeseAndMushroom.getName(), pizza1WithExtraCheeseAndMushroom.cost())

	pizza2 := NewFarmHouse()

	printPrice(pizza2.getName(), pizza2.cost())

	pizza2WithExtraCheese := NewExtraCheese(pizza2)

	printPrice(pizza2WithExtraCheese.getName(), pizza2WithExtraCheese.cost())

	pizza2WithExtraCheeseAndMushroom := NewMushroom(pizza2WithExtraCheese)

	printPrice(pizza2WithExtraCheeseAndMushroom.getName(), pizza2WithExtraCheeseAndMushroom.cost())

}

func printPrice(pizzaName string, price int) {
	fmt.Printf("Cost of %s is %d\n", pizzaName, price)
}
