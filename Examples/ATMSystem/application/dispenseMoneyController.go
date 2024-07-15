package application

import (
	"atm/models"
	"errors"
)

type DispenseMoney struct {
	State
}

func (d *DispenseMoney) InsertCard(machine *ATM, card *models.Card) (string, error) {
	return "", errors.New("Cash withdrawal is in progress")
}

func (d *DispenseMoney) EnterPin(machine *ATM, pin string) (string, error) {
	return "", errors.New("Cash withdrawal is in progress")
}

func (d *DispenseMoney) ViewBalance(machine *ATM) (string, error) {
	return "", errors.New("Cash withdrawal is in progress")
}

func (d *DispenseMoney) WithdrawCash(machine *ATM, amount float64) (string, error) {
	return "", errors.New("Cash withdrawal is in progress")
}

func (d *DispenseMoney) Cancel(machine *ATM) (string, error) {
	return "", errors.New("Cash withdrawal is in progress")
}
