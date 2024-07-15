package application

import (
	"errors"
	"fmt"
)

type IWithdrawProcessor interface {
	ProcessAmount(amount float64) (string, error)
}

type WithdrawProcessor struct {
	nextWithdrawProcessor IWithdrawProcessor
}

func NewWithDrawPeocessor() *WithdrawProcessor {
	return &WithdrawProcessor{
		nextWithdrawProcessor: &TwoThousandProcessor{
			nextWithdrawProcessor: &OneThousandProcessor{
				nextWithdrawProcessor: &FiveHundredProcessor{
					nextWithdrawProcessor: &TwoHandredProcessor{
						nextWithdrawProcessor: &HundredProcessor{
							nextWithdrawProcessor: nil,
						},
					},
				},
			},
		},
	}
}

func (p *WithdrawProcessor) ProcessAmount(amount float64) (string, error) {
	return p.nextWithdrawProcessor.ProcessAmount(amount)
}

type TwoThousandProcessor struct {
	nextWithdrawProcessor IWithdrawProcessor
}

func (p *TwoThousandProcessor) ProcessAmount(amount float64) (string, error) {
	noteCount := int64(amount) / 2000
	remainingAmount := amount - float64(2000*noteCount)
	if remainingAmount > 0 {
		resp, err := p.nextWithdrawProcessor.ProcessAmount(remainingAmount)
		if err == nil {
			if noteCount > 0 {
				return fmt.Sprintf("2000: %d,\n %s", noteCount, resp), nil
			}
			return resp, nil
		} else {
			return "", err
		}
	}
	return fmt.Sprintf("2000: %d\n", noteCount), nil
}

type OneThousandProcessor struct {
	nextWithdrawProcessor IWithdrawProcessor
}

func (p *OneThousandProcessor) ProcessAmount(amount float64) (string, error) {
	noteCount := int64(amount) / 1000
	remainingAmount := amount - float64(1000*noteCount)
	if remainingAmount > 0 {
		resp, err := p.nextWithdrawProcessor.ProcessAmount(remainingAmount)
		if err == nil {
			if noteCount > 0 {
				return fmt.Sprintf("1000: %d,\n %s", noteCount, resp), nil
			}
			return resp, nil
		} else {
			return "", err
		}
	}
	return fmt.Sprintf("1000: %d\n", noteCount), nil
}

type FiveHundredProcessor struct {
	nextWithdrawProcessor IWithdrawProcessor
}

func (p *FiveHundredProcessor) ProcessAmount(amount float64) (string, error) {
	noteCount := int64(amount) / 500
	remainingAmount := amount - float64(500*noteCount)
	if remainingAmount > 0 {
		resp, err := p.nextWithdrawProcessor.ProcessAmount(remainingAmount)
		if err == nil {
			if noteCount > 0 {
				return fmt.Sprintf("500: %d,\n %s", noteCount, resp), nil
			}
			return resp, nil
		} else {
			return "", err
		}
	}
	return fmt.Sprintf("500: %d\n", noteCount), nil
}

type TwoHandredProcessor struct {
	nextWithdrawProcessor IWithdrawProcessor
}

func (p *TwoHandredProcessor) ProcessAmount(amount float64) (string, error) {
	noteCount := int64(amount) / 200
	remainingAmount := amount - float64(500*noteCount)
	if remainingAmount > 0 {
		resp, err := p.nextWithdrawProcessor.ProcessAmount(remainingAmount)
		if err == nil {
			if noteCount > 0 {
				return fmt.Sprintf("200: %d,\n %s", noteCount, resp), nil
			}
			return resp, nil
		} else {
			return "", err
		}
	}
	return fmt.Sprintf("200: %d\n", noteCount), nil
}

type HundredProcessor struct {
	nextWithdrawProcessor IWithdrawProcessor
}

func (p *HundredProcessor) ProcessAmount(amount float64) (string, error) {
	noteCount := int64(amount) / 100
	remainingAmount := amount - float64(100*noteCount)
	if remainingAmount > 0 {
		return "", errors.New("amount should be multiple of 100")
	}
	return fmt.Sprintf("100: %d\n", noteCount), nil
}
