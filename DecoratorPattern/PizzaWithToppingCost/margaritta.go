package main

type Margaritta struct {
	name string
}

func NewMargaritta() *Margaritta {
	return &Margaritta{name: "Margaritta Base Pizza"}
}

func (m *Margaritta) cost() int {
	return 100
}

func (m *Margaritta) getName() string {
	return m.name
}
