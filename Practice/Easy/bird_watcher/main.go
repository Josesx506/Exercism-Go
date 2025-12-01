package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")
	total := 0
	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
	start := 7 * (week - 1)
	end := 7 * week
	weeklyBirds := birdsPerDay[start:end]
	return TotalBirdCount(weeklyBirds)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	// => 12
}
