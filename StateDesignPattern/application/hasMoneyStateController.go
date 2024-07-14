package application

import (
	"errors"
)

type HasMoney struct {
	State
}

func (d *HasMoney) InsertCoin(coin Coin, machine *VendingMachine) (string, error) {
	return machine.AddCoin(coin)
}

func (d *HasMoney) SelectProduct(productCode int, machine *VendingMachine) (string, error) {
	resp, err := machine.PickProduct(productCode)
	if err == nil {
		machine.UpdateState(&ProductSelection{})
	}
	return resp, err
}

func (d *HasMoney) DispenseProduct(machine *VendingMachine) (string, error) {
	return "", errors.New("Select Product first.")
}

func (d *HasMoney) Cancel(machine *VendingMachine) (string, error) {
	resp := machine.Cancel()
	machine.UpdateState(&IDLE{})
	return resp, nil
}
