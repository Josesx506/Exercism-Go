package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	// panic("Please implement the Convert function")
	output := ""
	if number%3 == 0 {
		output += "Pling"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		output += strconv.Itoa(number)
	}
	return output
}

func main() {
	fmt.Println(Convert(6))
	fmt.Println(Convert(13))
	fmt.Println(Convert(15))
	fmt.Println(Convert(21))
	fmt.Println(Convert(25))
}
