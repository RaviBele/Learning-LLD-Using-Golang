package application

type State interface {
	InsertCoin(coin Coin, machine *VendingMachine) (string, error)
	SelectProduct(productCode int, machine *VendingMachine) (string, error)
	DispenseProduct(machine *VendingMachine) (string, error)
	Cancel(machine *VendingMachine) (string, error)
}
