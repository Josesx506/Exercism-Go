package main

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	// panic("RemainingOvenTime not implemented")
	remaining := OvenTime - actualMinutesInOven
	return remaining
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	// panic("PreparationTime not implemented")
	prepTime := numberOfLayers * 2
	return prepTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// panic("ElapsedTime not implemented")
	prepTime := PreparationTime(numberOfLayers)
	elapsedTime := prepTime + actualMinutesInOven
	return elapsedTime
}

func main() {
	println(ElapsedTime(3, 15))
	println(RemainingOvenTime(10))
}
