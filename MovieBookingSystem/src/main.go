package main

import (
	"fmt"
	"math/rand"
	"moviebookingsystem/application"
	"time"
)

func main() {
	bookMyShow := application.BookMyShowApplication()

	pvr := bookMyShow.AddTheatreToCity("bangalore", "PVR")

	// imax := bookMyShow.AddTheatreToCity("bangalore", "imax")

	deadpoool := bookMyShow.AddMovie("bangalore", "Deadpool", 2.5)

	bahubali := bookMyShow.AddMovie("bangalore", "Bahubali", 3.0)

	pvrScreen1 := bookMyShow.AddScreen(pvr, 400, 20, 20, 20, 1)

	pvrScreen2 := bookMyShow.AddScreen(pvr, 400, 20, 20, 20, 2)

	// imaxScreen1 := bookMyShow.AddScreen(imax, 600, 40, 30, 20, 1)

	// imaxScreen2 := bookMyShow.AddScreen(imax, 600, 40, 30, 20, 2)

	pvrDeadPoolShow := bookMyShow.AddShow(pvr, time.Now().Local().Add(time.Hour*time.Duration(1)), pvrScreen1, deadpoool)
	pvrBahuBaliShow := bookMyShow.AddShow(pvr, time.Now().Local().Add(time.Hour*time.Duration(1)), pvrScreen2, bahubali)

	// imaxDeadPoolShow := bookMyShow.AddShow(imax, time.Now().Local().Add(time.Minute*time.Duration(30)), imaxScreen1, deadpoool)
	// imaxBahuBaliShow := bookMyShow.AddShow(imax, time.Now().Local().Add(time.Minute*time.Duration(30)), imaxScreen2, bahubali)

	deadPoolShowSeats := bookMyShow.GetAvailableSeatsForShow(pvr, pvrDeadPoolShow)

	startSeat := rand.Intn(len(deadPoolShowSeats)-5) + 5
	booking := bookMyShow.MakeBooking(pvr, pvrBahuBaliShow, deadPoolShowSeats[startSeat:startSeat+5])

	fmt.Printf("Booking: \n%v\n", booking)
}
