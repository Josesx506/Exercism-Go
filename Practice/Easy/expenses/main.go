package main

import (
	"errors"
	"fmt"
)

// Learn about first class functions & anonymous functions plus how they're used in go

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	// panic("Please implement the Filter function")
	filtered := make([]Record, 0, len(in))
	for _, record := range in {
		if predicate(record) {
			filtered = append(filtered, record)
		}
	}
	return filtered
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	// panic("Please implement the ByDaysPeriod function")
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	// panic("Please implement the ByCategory function")
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	// panic("Please implement the TotalByPeriod function")
	valid := Filter(in, ByDaysPeriod(p))
	total := 0.0
	for _, rec := range valid {
		total += rec.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	// panic("Please implement the CategoryExpenses function")
	categoryRecords := Filter(in, ByCategory(c))
	if len(categoryRecords) > 0 {
		periodRecords := Filter(categoryRecords, ByDaysPeriod(p))
		if len(periodRecords) > 0 {
			total := 0.0
			for _, rec := range periodRecords {
				total += rec.Amount
			}
			return total, nil
		} else {
			return 0.0, nil
		}
	} else {
		// An error must be returned only if there are no records in the list that
		// belong to the given category, regardless of period of time.
		return 0.0, errors.New("No records found")
	}
}

func main() {
	records := []Record{
		{
			Day:      1,
			Amount:   5.15,
			Category: "groceries",
		},
		{
			Day:      1,
			Amount:   3.45,
			Category: "groceries",
		},
		{
			Day:      13,
			Amount:   55.67,
			Category: "utility-bills",
		},
		{
			Day:      15,
			Amount:   11,
			Category: "groceries",
		},
	}
	period := DaysPeriod{
		From: 1,
		To:   15,
	}

	outRecords := Filter(records, ByDaysPeriod(period))
	fmt.Println(outRecords)
	outCategories := Filter(records, ByCategory("groceries"))
	fmt.Println(outCategories)
}
