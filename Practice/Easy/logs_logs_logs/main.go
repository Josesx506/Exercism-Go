package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
learn about runes and string  manipulation
*/

// Application identifies the application emitting the given log.
func Application(log string) string {
	// panic("Please implement the Application() function")
	for _, val := range log {
		switch { // omit the default case to prevent always triggering early return
		case string(val) == "❗":
			return "recommendation"
		case string(val) == "🔍":
			return "search"
		case string(val) == "☀":
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// panic("Please implement the Replace() function")
	updated := strings.ReplaceAll(log, string(oldRune), string(newRune))
	return updated
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// panic("Please implement the WithinLimit() function")
	count := utf8.RuneCountInString(log)
	return count <= limit
}

func main() {
	fmt.Println(Application("❗ recommended product"))
	fmt.Println(Application("executed search 🔍"))
	fmt.Println(Application("forecast: ☀ sunny"))
	fmt.Println(Application("🔍 search recommended product ❗"))
	fmt.Println(Replace("❗ recommended product", '❗', '?'))
	fmt.Println(WithinLimit("exercism❗", 9))
	fmt.Println(WithinLimit("exercism❗", 10))
	fmt.Println(WithinLimit("exercism❗", 8))
}
