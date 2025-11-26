package main

import (
	"fmt"
)

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	// panic("Please implement the FavoriteCards function")
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	// panic("Please implement the GetItem function")
	if index < 0 || index > len(slice)-1 {
		return -1
	} else {
		card := slice[index]
		return card
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	// panic("Please implement the SetItem function")
	if index > -1 && index < len(slice) {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// panic("Please implement the PrependItems function")
	size := len(slice) + len(values)
	out := make([]int, 0, size)
	out = append(out, values...)
	out = append(out, slice...)
	return out
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// panic("Please implement the RemoveItem function")
	if index > -1 && index < len(slice) {
		// remove
		out := append(slice[:index], slice[index+1:]...)
		return out
	} else {
		return slice
	}
}

func main() {
	fmt.Println(RemoveItem([]int{3, 4, 5, 6}, 3))
	fmt.Println(PrependItems([]int{3, 4, 5, 6}, -1, -2))
}
