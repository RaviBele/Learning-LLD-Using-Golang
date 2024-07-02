package main

type VegDelight struct {
	name string
}

func NewVegDelight() *VegDelight {
	return &VegDelight{name: "VegDelight Base Pizza"}
}

func (m *VegDelight) cost() int {
	return 150
}

func (m *VegDelight) getName() string {
	return m.name
}
