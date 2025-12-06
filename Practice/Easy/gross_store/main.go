package main

import "fmt"

/**
learn about maps and conditional statements in go
*/

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function")
	unit := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// panic("Please implement the NewBill() function")
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
	value, exists := units[unit]

	if exists {
		bill[item] += value
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the RemoveItem() function")
	unitValue, unitExists := units[unit]
	billValue, billExists := bill[item]

	if unitExists && billExists {
		switch {
		case billValue-unitValue < 0:
			return false
		case billValue-unitValue == 0:
			delete(bill, item)
			return true
		default:
			bill[item] -= unitValue
			return true
		}
	} else {
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// panic("Please implement the GetItem() function")
	value, exists := bill[item]
	return value, exists
}

func main() {
	units := Units()
	bill := NewBill()
	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)                                                   //true
	fmt.Println(AddItem(bill, units, "carrot", "gross"))              // true
	fmt.Println(RemoveItem(bill, units, "apples", "gross"))           // false
	fmt.Println(RemoveItem(bill, units, "carrot", "half_of_a_dozen")) // true
	val, _ := GetItem(bill, "carrot")
	fmt.Printf("%d carrots are on the customers bill\n", val)
}
