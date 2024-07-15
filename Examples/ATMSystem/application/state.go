package application

import "atm/models"

type State interface {
	InsertCard(machine *ATM, card *models.Card) (string, error)
	EnterPin(machine *ATM, pin string) (string, error)
	ViewBalance(machine *ATM) (string, error)
	WithdrawCash(machine *ATM, amount float64) (string, error)
	Cancel(machine *ATM) (string, error)
}
