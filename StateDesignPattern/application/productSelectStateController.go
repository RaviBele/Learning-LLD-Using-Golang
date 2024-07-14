package application

import (
	"errors"
)

type ProductSelection struct {
	State
}

func (d *ProductSelection) InsertCoin(coin Coin, machine *VendingMachine) (string, error) {
	return "", errors.New("Cannot insert coin now, Product selected, want for dispense of product?")
}

func (d *ProductSelection) SelectProduct(productCode int, machine *VendingMachine) (string, error) {
	return "", errors.New("Product already selected, want for dispense of product?")
}

func (d *ProductSelection) DispenseProduct(machine *VendingMachine) (string, error) {
	resp := machine.DispenseProduct()
	machine.UpdateState(&IDLE{})
	return resp, nil
}

func (d *ProductSelection) Cancel(machine *VendingMachine) (string, error) {
	return "", errors.New("Cannot cancel now, dispensing the product.")
}
