package main

import (
	"fmt"
	"time"
)

/**
Learn about parsing and formatting time objects in Go
*/

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// panic("Please implement the Schedule function")
	format := "1/2/2006 15:04:05"
	appt, _ := time.Parse(format, date)
	return appt
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// panic("Please implement the HasPassed function")
	format := "January 2, 2006 15:04:05"
	appt, _ := time.Parse(format, date)
	return time.Since(appt) > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// panic("Please implement the IsAfternoonAppointment function")
	format := "Monday, January 2, 2006 15:04:05"
	appt, _ := time.Parse(format, date)
	return appt.Hour() >= 12 && appt.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// panic("Please implement the Description function")
	format := "1/2/2006 15:04:05"
	ti, _ := time.Parse(format, date)
	out := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", ti.Weekday(), ti.Month().String(), ti.Day(), ti.Year(), ti.Hour(), ti.Minute())
	return out
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// panic("Please implement the AnniversaryDate function")
	// format := "2006-01-02 15:04:05 +0000"
	t := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}

func main() {
	fmt.Println(Schedule("12/6/2025 12:17:00"))
	fmt.Println(Description("12/6/2025 12:17:00"))
}
