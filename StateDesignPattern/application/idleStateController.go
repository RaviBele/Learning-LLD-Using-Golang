package application

import (
	"errors"
)

type IDLE struct {
	State
}

func (d *IDLE) InsertCoin(coin Coin, machine *VendingMachine) (string, error) {
	machine.UpdateState(&HasMoney{})
	return machine.AddCoin(coin)
}

func (d *IDLE) SelectProduct(productCode int, machine *VendingMachine) (string, error) {
	return "", errors.New("Insert coins first.")
}

func (d *IDLE) DispenseProduct(machine *VendingMachine) (string, error) {
	return "", errors.New("Insert coins first.")
}

func (d *IDLE) Cancel(machine *VendingMachine) (string, error) {
	return "", nil
}
