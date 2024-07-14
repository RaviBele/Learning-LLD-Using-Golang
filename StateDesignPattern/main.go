package main

import (
	"fmt"
	"statedesign/application"
)

func main() {
	product1 := application.NewProduct("COLA", 101, 20, 10)

	product2 := application.NewProduct("LIMCA", 102, 25, 10)

	product3 := application.NewProduct("FANTA", 103, 15, 10)

	vendingMachine := &application.VendingMachine{
		CurrentState: &application.IDLE{},
		Products: map[int]*application.Product{
			product1.GetCode(): product1,
			product2.GetCode(): product2,
			product3.GetCode(): product3,
		},
	}

	response, err := vendingMachine.InsertCoin(5)
	if err != nil {
		fmt.Printf("Error selecting product %s: %s\n", product1.GetName(), err.Error())
	}
	fmt.Println(response)

	response, err = vendingMachine.InsertCoin(5)
	if err != nil {
		fmt.Printf("Error selecting product %s: %s\n", product1.GetName(), err.Error())
	}
	fmt.Println(response)

	response, err = vendingMachine.InsertCoin(5)
	if err != nil {
		fmt.Printf("Error selecting product %s: %s\n", product1.GetName(), err.Error())
	}
	fmt.Println(response)

	response, err = vendingMachine.InsertCoin(5)
	if err != nil {
		fmt.Printf("Error selecting product %s: %s\n", product1.GetName(), err.Error())
	}
	fmt.Println(response)
	// response, err = vendingMachine.CancelButton()
	// if err != nil {
	// 	fmt.Printf("Error selecting product %s: %s\n", product3.GetName(), err.Error())
	// }
	// fmt.Println(response)
	response, err = vendingMachine.SelectProduct(product1.GetCode())
	if err != nil {
		fmt.Printf("Error selecting product %s: %s\n", product1.GetName(), err.Error())
	}
	fmt.Println(response)
	response, err = vendingMachine.DispenseSelectedProduct()
	if err != nil {
		fmt.Printf("Error selecting product %s: %s\n", product1.GetName(), err.Error())
	}
	fmt.Println(response)
}
