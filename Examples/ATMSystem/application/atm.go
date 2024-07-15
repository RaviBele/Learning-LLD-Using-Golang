package application

import (
	"atm/models"
	"errors"
	"fmt"
)

type ATM struct {
	currentState       State
	witbhdrawProcessor IWithdrawProcessor
	currentCard        *models.Card
	options            []string
	pinRetryCount      int
}

func NewATM() *ATM {
	return &ATM{
		currentState:       &IDLE{},
		witbhdrawProcessor: NewWithDrawPeocessor(),
		pinRetryCount:      3,
		options:            []string{"View Balance", "Cash Withdraw", "Cancel"},
	}
}

func (m *ATM) UpdateState(state State) {
	m.currentState = state
}

func (m *ATM) InsertCard(card *models.Card) (string, error) {
	return m.currentState.InsertCard(m, card)
}

func (m *ATM) EnterPin(pin string) (string, error) {
	return m.currentState.EnterPin(m, pin)
}

func (m *ATM) ViewBalance() (string, error) {
	return m.currentState.ViewBalance(m)
}

func (m *ATM) WithdrawCash(amount float64) (string, error) {
	resp, err := m.currentState.WithdrawCash(m, amount)
	return resp, err
}

func (m *ATM) Cancel() (string, error) {
	return m.currentState.Cancel(m)
}

func (m *ATM) AcceptCard(card *models.Card) (string, error) {
	m.currentCard = card
	return "Card accepted, Insert your PIN", nil
}

func (m *ATM) VerifyPin(pin string) (string, error) {
	if m.pinRetryCount <= 0 {
		return "", errors.New("you reached the limit for your PIN")
	}
	if m.currentCard.GetPin() == pin {
		m.pinRetryCount = 3
		return fmt.Sprintf("PIN is correct\n Please select below options\n %v", m.options), nil
	}
	m.pinRetryCount -= 1 //
	return "", errors.New("invalid PIN, try again.")
}

func (m *ATM) ReturnCard() (string, error) {
	m.currentCard = nil
	m.pinRetryCount = 3
	return "Please collect your card", nil
}

func (m *ATM) GetBalance() (string, error) {
	return fmt.Sprintf("%f", m.currentCard.GetAccount().GetBalance()), nil
}

func (m *ATM) WithdrawRequestCash(amount float64) (string, error) {
	currentBalance := m.currentCard.GetAccount().GetBalance()

	if currentBalance < amount {
		return "", errors.New("insuffienct balance")
	}

	resp, err := m.witbhdrawProcessor.ProcessAmount(amount)
	if err == nil {
		m.currentCard.GetAccount().UpdateBalance(currentBalance - amount)
	}
	return resp, err
}
