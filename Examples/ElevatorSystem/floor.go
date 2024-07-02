package main

type Floor struct {
	id             int
	display        *Display
	externalButton *ExternalButton
}

func NewFloor(id int) *Floor {
	return &Floor{
		id:             id,
		display:        NewDisplay(),
		externalButton: NewExternalButton(),
	}
}

func (f *Floor) pressButton(direction Direction) {
	f.externalButton.press(f.id, direction)
}

func (f *Floor) setDisplay(floor int, direction Direction) {
	f.display.setDisplayMessage(floor, direction)
}

func (f *Floor) getID() int {
	return f.id
}

func (f *Floor) Button() *ExternalButton {
	return f.externalButton
}
