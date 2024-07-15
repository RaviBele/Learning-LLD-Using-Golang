package application

import (
	"atm/models"
	"errors"
)

type HasCard struct {
	State
}

func (d *HasCard) InsertCard(machine *ATM, card *models.Card) (string, error) {
	return "", errors.New("Cannot insert more than one card")
}

func (d *HasCard) EnterPin(machine *ATM, pin string) (string, error) {
	resp, err := machine.VerifyPin(pin)
	if err == nil {
		machine.UpdateState(&SelectOptions{})
	}
	return resp, err
}

func (d *HasCard) ViewBalance(machine *ATM) (string, error) {
	return "", errors.New("Insert PIN First")
}

func (d *HasCard) WithdrawCash(machine *ATM, amount float64) (string, error) {
	return "", errors.New("insert PIN first")
}

func (d *HasCard) Cancel(machine *ATM) (string, error) {
	resp, err := machine.ReturnCard()
	if err == nil {
		machine.UpdateState(&IDLE{})
	}
	return resp, err
}
