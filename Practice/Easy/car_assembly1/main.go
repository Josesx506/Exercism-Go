package main

/**
learn about int and float variable types
*/

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	perHour := float64(productionRate) * (successRate / 100)
	return perHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	perMinute := float64(productionRate/60) * (successRate / 100)
	return int(perMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	discountPrice10 := int(carsCount/10) * 95000
	singleCarPrice := (carsCount % 10) * 10000
	return uint(discountPrice10 + singleCarPrice)
}

func main() {
	println(CalculateWorkingCarsPerHour(221, 90))
	println(CalculateWorkingCarsPerMinute(221, 95))
	println(CalculateCost(12))
}
