package main

type Admin struct {
	Account
}

func NewAdmin(username string) *Admin {
	return &Admin{
		Account: Account{
			username: username,
		},
	}
}

func (a *Admin) addParkingFloor(parking *ParkingFloor) {
	getParkingLot().addParkingFloor(parking)
}

func (a *Admin) addEntryPanel(entryPanel *EntryPanel) {
	getParkingLot().addEntryPanel(entryPanel)
}

func (a *Admin) addExitPanel(exitPanel *ExitPanel) {
	getParkingLot().addExitPanel(exitPanel)
}

func (a *Admin) addParkingSpot(floorID string, spot IParkingSpot) {
	floors := getParkingLot().getListOfParkingFloors()

	for _, floor := range floors {
		if floor.getParkingFloorID() == floorID {
			floor.addParkingSpot(spot)
			return
		}
	}
}
