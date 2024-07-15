package application

import (
	"atm/models"
	"errors"
	"fmt"
	"time"
)

type SelectOptions struct {
	State
}

func (d *SelectOptions) InsertCard(machine *ATM, card *models.Card) (string, error) {
	return "", errors.New("Cannot insert more than one card")
}

func (d *SelectOptions) EnterPin(machine *ATM, pin string) (string, error) {
	return "", errors.New("Already inserted PIN")
}

func (d *SelectOptions) ViewBalance(machine *ATM) (string, error) {
	return machine.GetBalance()
}

func (d *SelectOptions) WithdrawCash(machine *ATM, amount float64) (string, error) {
	resp, err := machine.WithdrawRequestCash(amount)
	if err == nil {
		machine.UpdateState(&DispenseMoney{})
		fmt.Println("Dispensing cash")
		time.Sleep(time.Second * 5)
		fmt.Println("Cash dispensed")
		machine.UpdateState(&SelectOptions{})
	}
	return resp, err
}

func (d *SelectOptions) Cancel(machine *ATM) (string, error) {
	resp, err := machine.ReturnCard()
	if err == nil {
		machine.UpdateState(&IDLE{})
	}
	return resp, err
}