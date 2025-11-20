package main

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
	res := fmt.Sprintf("Welcome to my party, %s!", name)
	return res
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
	res := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return res
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")
	res := Welcome(name) + "\n"
	res = res + fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
	res = res + fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return res
}

func main() {
	println(Welcome("Xuân Jing"))
	println(HappyBirthday("Xuân Jing", 17))
	println(AssignTable("Paula", 101, "Chioma", "on the right", 100.0))
}
