// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package main

import "fmt"

// learn about string manipulation

// ShareWith returns a string describing how cookies are
// shared between 2 individuals.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}

func main() {
	fmt.Println(ShareWith("John"))
	fmt.Println(ShareWith("Alice"))
	fmt.Println(ShareWith("Bahdun"))
	fmt.Println(ShareWith(""))
}
