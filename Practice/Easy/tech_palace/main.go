package main

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
	result := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return result
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
	stars := strings.Repeat("*", numStarsPerLine)
	result := stars + "\n" + welcomeMsg + "\n" + stars
	return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
	clean := strings.ReplaceAll(oldMsg, "*", "")
	clean = strings.TrimSpace(clean)

	return clean
}

func main() {
	println(WelcomeMessage("Judy"))
	println(AddBorder("Welcome to tech palace", 25))
	println(CleanupMessage("**************************\n*    BUY NOW, SAVE 10%   *\n**************************"))
}
