package main

import (
	"atm/application"
	"atm/models"
	"fmt"
)

func main() {
	atm := application.NewATM()

	user := models.NewUser("Ravi")
	account := models.NewAccount(user.GetName(), 1000)
	card := models.NewCard(account)
	user.SetCard(card)
	user.SetAccount(account)

	resp, err := atm.InsertCard(card)
	if err != nil {
		fmt.Println("Failed to insert card: ", err.Error())
	}
	fmt.Println(resp, "\n")

	resp, err = atm.EnterPin(card.GetPin())
	if err != nil {
		fmt.Println("Failed in PIN: ", err.Error())
	}
	fmt.Println(resp, "\n")

	fmt.Println("Selected view balance\n")
	resp, err = atm.ViewBalance()
	if err != nil {
		fmt.Println("Failed to account Balance: ", err.Error())
	}
	fmt.Println(resp, "\n")

	fmt.Println("Withdrawing amount: 3900\n")
	resp, err = atm.WithdrawCash(3900)
	if err != nil {
		fmt.Println("Failed to withdraw amount: ", err.Error())
	} else {
		fmt.Println("\nHere is you cash\n")
	}
	fmt.Println(resp, "\n")

	fmt.Println("Selected view balance\n")
	resp, err = atm.ViewBalance()
	if err != nil {
		fmt.Println("Failed to account Balance: ", err.Error())
	}
	fmt.Println(resp, "\n")

	fmt.Println("Canceling...\n")
	resp, err = atm.Cancel()
	if err != nil {
		fmt.Println("Failed to cancel: ", err.Error())
	}
	fmt.Println(resp, "\n")

	fmt.Println("Selected view balance\n")
	resp, err = atm.ViewBalance()
	if err != nil {
		fmt.Println("Failed to account Balance: ", err.Error())
	}
	fmt.Println(resp, "\n")
}
