package main

import (
	"fmt"
	"regexp"
)

type File [8]bool

type Chessboard map[string]*File

func NewChessboard() Chessboard {
	cb := Chessboard{
		"A": &File{true, false, true, false, true, false, true, false},
		"B": &File{true, false, true, false, true, false, true, false},
		"C": &File{true, false, true, false, true, false, true, false},
		"D": &File{true, false, true, false, true, false, true, false},
		"E": &File{true, false, true, false, true, false, true, false},
		"F": &File{true, false, true, false, true, false, true, false},
		"G": &File{true, false, true, false, true, false, true, false},
		"H": &File{true, false, true, false, true, false, true, false},
	}

	return cb
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func TagWithUserName(lines []string) []string {
	// panic("Please implement the TagWithUserName function")
	re, _ := regexp.Compile(`User\s+(\S+)`)
	for i, text := range lines {
		sl := re.FindStringSubmatch(text)
		if len(sl) > 1 {
			lines[i] = "[USR] " + sl[1] + " " + text
			fmt.Println("[USR] " + sl[1] + " " + text + "\n")
		}
	}
	return lines
}

func main() {
	cb := NewChessboard()
	// cb := Chessboard{
	// 	"A": File{true, false, true, false, true, false, true, false},
	// 	"B": File{true, false, true, false, true, false, true, false},
	// 	"C": File{true, false, true, false, true, false, true, false},
	// 	"D": File{true, false, true, false, true, false, true, false},
	// 	"E": File{true, false, true, false, true, false, true, false},
	// 	"F": File{true, false, true, false, true, false, true, false},
	// 	"G": File{true, false, true, false, true, false, true, false},
	// 	"H": File{true, false, true, false, true, false, true, false},
	// }

	println(cb)

	cb["B"][0] = false
	println(cb["B"][0])

	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	lines := []string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	}
	TagWithUserName(lines)
	// fmt.Println()
}
