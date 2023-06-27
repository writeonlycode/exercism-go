package speed

// Struct Car represents a car with its current battery level percentage, the
// battery drain rate, the speed of the car and the current distance traveled.
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given
// specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, batteryDrain: batteryDrain, speed: speed, distance: 0}
}

// Struct Track represents a track with a given distance in meters.
type Track struct {
	distance int
}

// NewTrack creates a new track with a given distance.
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one
// more time, the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}

	car.battery -= car.batteryDrain
	car.distance += car.speed

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	numberOfDrives := track.distance / car.speed
	batteryNeeded := numberOfDrives * car.batteryDrain
	return car.battery >= batteryNeeded
}
