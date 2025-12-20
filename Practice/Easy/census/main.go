package main

import "fmt"

// Learn about zero values for different variable types

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	// panic("Please implement NewResident.")
	resident := &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}

	return resident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	// panic("Please implement HasRequiredInfo.")
	street, streetOk := r.Address["street"]
	return r.Name != "" && streetOk && street != ""
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	// panic("Please implement Delete.")
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	// panic("Please implement Count.")
	total := 0

	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			total++
		}
	}
	return total
}

func main() {
	validResident := NewResident("Matthew Sanabria", 29,
		map[string]string{"street": "Main St."})
	invalidResident := NewResident("", 29,
		map[string]string{"street": "Main St."})

	fmt.Println(validResident)
	fmt.Println(validResident.HasRequiredInfo())
	fmt.Println(invalidResident.HasRequiredInfo())

	invalidResident.Delete()
	fmt.Println(invalidResident)

	residents := make([]*Resident, 0, 10)
	residents = append(residents, validResident)
	fmt.Println(Count(residents))
}
