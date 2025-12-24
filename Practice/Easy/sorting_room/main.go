package main

import (
	"fmt"
	"strconv"
)

// Learn about type assertions and conversions

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	// panic("Please implement DescribeNumber")
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	// panic("Please implement DescribeNumberBox")
	num := nb.Number()
	fnum := float64(num)
	return fmt.Sprintf("This is a box containing the number %.1f", fnum)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	// panic("Please implement ExtractFancyNumber")
	str, ok := fnb.(FancyNumber)
	if ok {
		val, _ := strconv.Atoi(str.Value())
		return val
	} else {
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	// panic("Please implement DescribeFancyNumberBox")
	val := ExtractFancyNumber(fnb)
	fval := float64(val)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", fval)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	// panic("Please implement DescribeAnything")
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}

func main() {
	var fnb FancyNumber = FancyNumber{
		n: "-5.2", // Any float breaks fancy num check
	}
	var fnb2 FancyNumber = FancyNumber{
		n: "10",
	}
	fmt.Println(DescribeNumber(4.1))
	fmt.Println(ExtractFancyNumber(fnb))
	fmt.Println(ExtractFancyNumber(fnb2))
	fmt.Println(DescribeFancyNumberBox(fnb))
	fmt.Println(DescribeAnything(fnb2))
	fmt.Println(DescribeAnything(12))
	fmt.Println(DescribeAnything(5.3))
}
