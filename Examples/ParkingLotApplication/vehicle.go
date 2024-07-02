package main

type VehicleType int

const (
	Car         VehicleType = iota + 1 // EnumIndex = 1
	Truck                              // EnumIndex = 2
	ElectricCar                        // EnumIndex = 3
	MotorBike                          // EnumIndex = 4
)

var validVehicleType = []VehicleType{Car, Truck, ElectricCar, MotorBike}

// String - Creating common behavior - give the type a String function
func (w VehicleType) String() string {
	return [...]string{"Car", "Truck", "ElectricCar", "MotorBike"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w VehicleType) EnumIndex() int {
	return int(w)
}

func isValidVehicleType(w VehicleType) bool {
	for _, v := range validVehicleType {
		if w == v {
			return true
		}
	}
	return false
}

type IVehicle interface {
	getVehicleType() VehicleType
	getVehicleNumber() string
	isWithDisabled() bool
}

type Vehicle struct {
	vehicleNumber string
	vehicleType   VehicleType
	withDisabled  bool
}

func (v *Vehicle) getVehicleType() VehicleType {
	return v.vehicleType
}

func (v *Vehicle) getVehicleNumber() string {
	return v.vehicleNumber
}

func (v *Vehicle) isWithDisabled() bool {
	return v.withDisabled
}
