package main

type FarmHouse struct {
	name string
}

func NewFarmHouse() *FarmHouse {
	return &FarmHouse{name: "FarmHouse Base Pizza"}
}

func (m *FarmHouse) cost() int {
	return 200
}

func (m *FarmHouse) getName() string {
	return m.name
}
