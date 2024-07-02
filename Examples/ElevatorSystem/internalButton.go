package main

type InternalButton struct {
	internalButtonDispatcher *internalButtonDispatcher
}

func NewInternalButton() *InternalButton {
	return &InternalButton{
		internalButtonDispatcher: NewInternalButtonDispatcher(),
	}
}

func (b *InternalButton) press(floorNumber int, dir Direction, elevatorID int) {
	b.internalButtonDispatcher.submitRequest(floorNumber, dir, elevatorID)
}

func (b *InternalButton) addController(controller *ElevatorController) {
	b.internalButtonDispatcher.addElevatorController(controller)
}
