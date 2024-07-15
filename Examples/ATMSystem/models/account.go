package models

import "atm/utils"

type Account struct {
	name    string
	number  string
	balance float64
}

func NewAccount(name string, balance float64) *Account {
	return &Account{
		name:    name,
		balance: balance,
		number:  utils.GenerateAccountNumber(),
	}
}

func (a *Account) GetName() string {
	return a.name
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) GetNumber() string {
	return a.number
}

func (a *Account) UpdateBalance(balance float64) {
	a.balance = balance
}
