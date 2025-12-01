package main

import "fmt"

/**
Learn to calculate values using the float32 and float64 values in go
*/

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	// panic("Please implement the InterestRate function")

	switch {
	case balance < 0:
		return float32(3.213)
	case balance >= 0 && balance < 1000:
		return float32(0.5)
	case balance >= 1000 && balance < 5000:
		return float32(1.621)
	default:
		return float32(2.475)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	// panic("Please implement the Interest function")
	interestVal := balance * float64(InterestRate(balance)/100.0)
	return interestVal
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	// panic("Please implement the AnnualBalanceUpdate function")
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// panic("Please implement the YearsBeforeDesiredBalance function")
	years := 0

	for balance < targetBalance {
		balance += Interest(balance)
		years++
	}

	return years
}

func main() {
	fmt.Println(YearsBeforeDesiredBalance(100.0, 125.80))
	fmt.Println(YearsBeforeDesiredBalance(1000.0, 1100.0))
	fmt.Println(YearsBeforeDesiredBalance(8080.80, 9090.90))
}
