package main

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	// panic("Please implement the NewCar function")
	newCar := Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
	return newCar
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	// panic("Please implement the NewTrack function")
	newTrack := Track{
		distance: distance,
	}
	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// panic("Please implement the Drive function")
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// panic("Please implement the CanFinish function")
	driveCycles := track.distance / car.speed
	if car.battery >= (driveCycles * car.batteryDrain) {
		return true
	} else {
		return false
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	distance := 100
	track := NewTrack(distance)

	println(CanFinish(car, track)) // true
}
