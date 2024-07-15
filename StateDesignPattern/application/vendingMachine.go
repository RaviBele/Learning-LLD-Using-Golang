package application

import (
	"errors"
	"fmt"
)

type VendingMachine struct {
	CurrentState       State
	Products           map[int]*Product
	currentMoney       int
	currentProductCode int
}

func (v *VendingMachine) UpdateState(state State) {
	v.CurrentState = state
}

func (v *VendingMachine) InsertCoin(coin Coin) (string, error) {
	return v.CurrentState.InsertCoin(coin, v)
}

func (v *VendingMachine) SelectProduct(productCode int) (string, error) {
	return v.CurrentState.SelectProduct(productCode, v)
}

func (v *VendingMachine) DispenseSelectedProduct() (string, error) {
	return v.CurrentState.DispenseProduct(v)
}

func (v *VendingMachine) CancelButton() (string, error) {
	return v.CurrentState.Cancel(v)
}

func (v *VendingMachine) AddCoin(coin Coin) (string, error) {
	if coin != ONE && coin != TWO && coin != FIVE && coin != TEN {
		return "", fmt.Errorf("invalid coin: %v", coin)
	}
	v.currentMoney += int(coin)
	return fmt.Sprintf("Coin %d inserted", coin), nil
}

func (v *VendingMachine) PickProduct(productCode int) (string, error) {
	productCount := v.Products[productCode].GetCount()
	if productCount == 0 {
		return "", errors.New("Product out of stock")
	}

	productPrice := v.Products[productCode].GetPrice()
	if v.currentMoney < productPrice {
		return "", fmt.Errorf("price %d, Not enough money: %d", productPrice, v.currentMoney)
	}

	v.currentProductCode = productCode
	return fmt.Sprintf("Product %d selected", productCode), nil
}

func (v *VendingMachine) DispenseProduct() string {
	productPrice := v.Products[v.currentProductCode].GetPrice()
	v.Products[v.currentProductCode].UpdateCount()
	changeMoney := 0
	if v.currentMoney > productPrice {
		changeMoney = v.currentMoney - productPrice
	}
	v.currentMoney = 0
	if changeMoney > 0 {
		return fmt.Sprintf("Here is your %s and your change: %d\n", v.Products[v.currentProductCode].GetName(), changeMoney)
	} else {
		return fmt.Sprintf("Here is your %s\n", v.Products[v.currentProductCode].GetName())
	}
}

func (v *VendingMachine) Cancel() string {
	resp := ""
	if v.currentMoney > 0 {
		resp = fmt.Sprintf("Canceling..., Here is your money back %d\n", v.currentMoney)
	} else {
		resp = "Canceling..."
	}
	v.currentMoney = 0

	return resp
}
