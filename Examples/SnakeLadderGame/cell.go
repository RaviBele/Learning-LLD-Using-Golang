package main

type Cell struct {
	jump *Jump
}

func NewCell(jump *Jump) *Cell {
	return &Cell{jump: jump}
}

func (c *Cell) hasJump() bool {
	return c.jump != nil
}

func (c *Cell) AssignJump(jump *Jump) {
	c.jump = jump
}

func (c *Cell) getJump() *Jump {
	return c.jump
}
