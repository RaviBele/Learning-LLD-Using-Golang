package models

import "atm/utils"

type Card struct {
	number  string
	pin     string
	account *Account
}

func NewCard(account *Account) *Card {
	return &Card{
		number:  utils.GenerateCardNumber(),
		pin:     utils.GeneratePINNumber(),
		account: account,
	}
}

func (c *Card) GetNumber() string {
	return c.number
}

func (c *Card) GetPin() string {
	return c.pin
}

func (c *Card) GetAccount() *Account {
	return c.account
}
