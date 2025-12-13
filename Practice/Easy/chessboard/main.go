package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	// panic("Please implement CountInFile()")
	count := 0
	for _, item := range cb[file] {
		if item {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	// panic("Please implement CountInRank()")
	count := 0
	if rank < 1 || rank > 8 {
		return count
	} else {
		for _, row := range cb {
			if row[rank-1] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// panic("Please implement CountAll()")
	count := 0
	for _, row := range cb {
		for range row {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	// panic("Please implement CountOccupied()")
	count := 0
	for _, row := range cb {
		for _, square := range row {
			if square {
				count++
			}
		}
	}
	return count
}

func NewChessboard() Chessboard {
	cb := Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{true, false, true, false, true, false, true, false},
		"C": File{true, false, true, false, true, false, true, false},
		"D": File{true, false, true, false, true, false, true, false},
		"E": File{true, false, true, false, true, false, true, false},
		"F": File{true, false, true, false, true, false, true, false},
		"G": File{true, false, true, false, true, false, true, false},
		"H": File{true, false, true, false, true, false, true, false},
	}

	return cb
}

func main() {
	fmt.Println("These are the tests.")
	cb := NewChessboard()
	fmt.Println(CountInFile(cb, "A")) // Count row A
	fmt.Println(CountInRank(cb, 3))   // Count column 3
	fmt.Println(CountAll(cb))         // All cells
	fmt.Println(CountOccupied(cb))    // All occuppied cells
}
