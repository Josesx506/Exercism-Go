package main

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var quanNoodles int
	var quanSauce float64

	for i := 0; i < len(layers); i++ {
		layer := layers[i]
		if layer == "noodles" {
			quanNoodles += 50
		} else if layer == "sauce" {
			quanSauce += 0.2
		}
	}

	return quanNoodles, quanSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	mult := float64(scale) / 2 // Each slice can originally produce 2 portions
	scaledQuants := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		scaledQuants[i] = quantities[i] * mult
	}

	return scaledQuants
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(Quantities(layers))

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)

	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)
}
