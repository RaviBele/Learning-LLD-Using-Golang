package main

type Building struct {
	floors []*Floor
}

func NewBuilding(noOfFloors int) *Building {
	building := &Building{}
	for i := 0; i < noOfFloors; i++ {
		building.floors = append(building.floors, NewFloor(i))
	}
	return building
}
