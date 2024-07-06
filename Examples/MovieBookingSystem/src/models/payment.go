package models

type Payment struct {
	id     string
	status string
}

func NewPayment() *Payment {
	return &Payment{status: "COMPLETED"}
}
