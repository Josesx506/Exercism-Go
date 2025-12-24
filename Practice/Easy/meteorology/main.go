package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}

func main() {
	ath := MeteorologyData{"Athens", Temperature{21, Celsius}, "N", Speed{16, KmPerHour}, 63}
	fmt.Println(ath)
	dlh := MeteorologyData{"Delhi", Temperature{33, Celsius}, "W", Speed{2, MilesPerHour}, 23}
	fmt.Println(dlh)
	sf := MeteorologyData{"San Francisco", Temperature{57, Fahrenheit}, "NW", Speed{19, MilesPerHour}, 60}
	fmt.Println(sf)
}
