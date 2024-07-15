package application

import (
	"atm/models"
	"errors"
)

type IDLE struct {
	State
}

func (d *IDLE) InsertCard(machine *ATM, card *models.Card) (string, error) {
	resp, err := machine.acceptCard(card)
	if err == nil {
		machine.updateState(&HasCard{})
	}

	return resp, err
}

func (d *IDLE) EnterPin(machine *ATM, pin string) (string, error) {
	return "", errors.New("insert card first")
}

func (d *IDLE) ViewBalance(machine *ATM) (string, error) {
	return "", errors.New("insert card first")
}

func (d *IDLE) WithdrawCash(machine *ATM, amount float64) (string, error) {
	return "", errors.New("insert card first")
}

func (d *IDLE) Cancel(machine *ATM) (string, error) {
	return "", nil
}
