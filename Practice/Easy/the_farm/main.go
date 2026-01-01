package main

import (
	"errors"
	"fmt"
)

type Fodder struct {
	amountPerCow float64
	factor       float64
}

func (f Fodder) FodderAmount(cows int) (float64, error) {
	if cows > 0 {
		return f.amountPerCow * float64(cows), nil
	} else {
		return 0, errors.New("Invalid number of cows")
	}
}
func (f Fodder) FatteningFactor() (float64, error) {
	if f.factor > 0 {
		return f.factor, nil
	} else {
		return f.factor, errors.New("Erroneous fattening factor")
	}
}

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nc int) (float64, error) {
	totalFodder, tfe := fc.FodderAmount(nc)
	if tfe != nil {
		return 0, tfe
	}
	fatFactor, ffe := fc.FatteningFactor()
	if ffe != nil {
		return 0, ffe
	}

	singleCowFodder := (totalFodder / float64(nc)) * fatFactor

	if singleCowFodder > 0 {
		return singleCowFodder, nil
	} else {
		return 0, errors.New("Single cow fodder value is less than 0")
	}
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nc int) (float64, error) {
	if nc > 0 {
		return DivideFood(fc, nc)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	cows    int
	message string
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.cows, ice.message)
}

func ValidateNumberOfCows(nc int) error {
	switch {
	case nc < 0:
		return &InvalidCowsError{
			cows: nc, message: "there are no negative cows",
		}
	case nc == 0:
		return &InvalidCowsError{
			cows: nc, message: "no cows don't need food",
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

func main() {
	fod := Fodder{
		amountPerCow: 10.0,
		factor:       1,
	}

	food, err := DivideFood(fod, 10)
	fmt.Println(food, err)

	fod.amountPerCow = 12.1
	fod.factor = 1.3
	food, err = DivideFood(fod, 5)
	fmt.Println(food, err)

	fod.amountPerCow = 0
	food, err = DivideFood(fod, 10)
	fmt.Println(food, err)

	fod.amountPerCow = 10.0
	fod.factor = 0
	food, err = DivideFood(fod, 10)
	fmt.Println(food, err)

	fod.factor = 1
	food, err = ValidateInputAndDivideFood(fod, 0)
	fmt.Println(food, err)
	food, err = ValidateInputAndDivideFood(fod, 10)
	fmt.Println(food, err)

	fmt.Println(ValidateNumberOfCows(-20))
	fmt.Println(ValidateNumberOfCows(0))
	fmt.Println(ValidateNumberOfCows(20))
}
