package main

// for snake and ladders
type Jump struct {
	start int
	end   int
}

func NewJump(start int, end int) *Jump {
	return &Jump{start: start, end: end}
}

func (j *Jump) getStart() int {
	return j.start
}

func (j *Jump) getEnd() int {
	return j.end
}
