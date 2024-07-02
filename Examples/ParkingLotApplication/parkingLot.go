package main

type ParkingLot struct {
	parkingFloors []*ParkingFloor
	entryPanels   []*EntryPanel
	exitPanels    []*ExitPanel
}

var parkingLot *ParkingLot

func getParkingLot() *ParkingLot {
	if parkingLot == nil {
		parkingLot = &ParkingLot{}
	}
	return parkingLot
}

func (l *ParkingLot) addParkingFloor(newFloor *ParkingFloor) {
	for _, floor := range l.parkingFloors {
		if floor.getParkingFloorID() == newFloor.getParkingFloorID() {
			return
		}
	}

	l.parkingFloors = append(l.parkingFloors, newFloor)
}

func (l *ParkingLot) addEntryPanel(newEntryPanel *EntryPanel) {
	for _, entryPanel := range l.entryPanels {
		if entryPanel.getID() == newEntryPanel.getID() {
			return
		}
	}

	l.entryPanels = append(l.entryPanels, newEntryPanel)
}

func (l *ParkingLot) addExitPanel(newExitPanel *ExitPanel) {
	for _, exitPanel := range l.exitPanels {
		if exitPanel.getID() == newExitPanel.getID() {
			return
		}
	}

	l.exitPanels = append(l.exitPanels, newExitPanel)
}

func (l *ParkingLot) getListOfParkingFloors() []*ParkingFloor {
	return l.parkingFloors
}

func (l *ParkingLot) getListOfEntryPanels() []*EntryPanel {
	return l.entryPanels
}

func (l *ParkingLot) getListOfExitPanels() []*ExitPanel {
	return l.exitPanels
}

func (l *ParkingLot) vacateParkingSpot(parkingSpotID string, parkingFloorID string, spotType ParkingSpotType) IParkingSpot {
	for _, floor := range l.parkingFloors {
		if floor.getParkingFloorID() == parkingFloorID {
			for _, spot := range floor.getListOfParkingSpots()[spotType] {
				if spot.getParkingSpotID() == parkingSpotID {
					spot.vacateSpot()
					return spot
				}
			}
		}
	}

	return nil
}
