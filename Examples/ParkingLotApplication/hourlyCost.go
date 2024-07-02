package main

import "fmt"

type HourlyCost struct {
	hourlyCost map[ParkingSpotType]int
}

var hourlyCost *HourlyCost

func getHourlyCostObject() *HourlyCost {
	if hourlyCost == nil {
		hourlyCost = &HourlyCost{
			hourlyCost: map[ParkingSpotType]int{
				Compact:         25,
				Disabled:        15,
				ElectricCarSpot: 35,
				Large:           45,
				MotorCycle:      15,
			},
		}
	}
	return hourlyCost
}

func (h *HourlyCost) getCost(parkingSpotType ParkingSpotType) (int, error) {
	_, ok := h.hourlyCost[parkingSpotType]

	if ok {
		return h.hourlyCost[parkingSpotType], nil
	}

	return 0, fmt.Errorf("Invalid parking spot type")
}
