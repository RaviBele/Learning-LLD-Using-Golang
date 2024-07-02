package main

type ExternalButton struct {
	externalButtonDispatcher *ExternalButtonDispatcher
}

func NewExternalButton() *ExternalButton {
	return &ExternalButton{
		externalButtonDispatcher: NewExternalButtonDispatcher(),
	}
}

func (b *ExternalButton) press(floor int, direction Direction) {
	b.externalButtonDispatcher.submitRequest(floor, direction)
}

func (b *ExternalButton) addController(controller *ElevatorController) {
	b.externalButtonDispatcher.addElevatorController(controller)
}
