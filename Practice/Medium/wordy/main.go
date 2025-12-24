package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Calculator(a, b int, opr string) int {
	switch opr {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func Answer(question string) (int, bool) {
	// panic("Please implement the Answer function")
	values := make([]int, 0)
	operators := make([]string, 0)

	re := regexp.MustCompile(`^What is (-?\d+)(?: (?:plus|minus|multiplied by|divided by) (?:-?\d+))*\?$`)
	validQ := re.FindAllStringSubmatch(question, -1)
	if validQ == nil {
		return 0, false
	}

	mapOp := map[string]string{
		"plus":       "+",
		"minus":      "-",
		"multiplied": "*",
		"divided":    "/",
	}

	question = strings.ReplaceAll(question, "?", "")
	words := strings.Split(question, " ")
	validOperators := []string{"multiplied", "plus", "minus", "divided"}

	for _, word := range words {
		if slices.Contains(validOperators, word) {
			operators = append(operators, mapOp[word])
		} else {
			num, err := strconv.Atoi(word)
			if err == nil {
				values = append(values, num)
			}
		}
	}

	ans := values[0]
	for idx, opr := range operators {
		ans = Calculator(ans, values[idx+1], opr)
	}
	return ans, true
}

// func Answer(question string) (int, bool) {
// 	// panic("Please implement the Answer function")
// 	values := make([]int, 0)
// 	operators := make([]string, 0)

// 	re := regexp.MustCompile(`^What is (-?\d+)(?: (?:plus|minus|multiplied by|divided by) (?:-?\d+))*\?$`)
// 	err := re.FindAllStringSubmatch(question, -1)
// 	if err == nil {
// 		return 0, false
// 	}

// 	question = strings.ReplaceAll(question, "?", "")
// 	words := strings.Split(question, " ")
// 	validOperators := []string{"multiplied", "plus", "minus", "divided"}

// 	for idx, word := range words {
// 		if idx > 0 && slices.Contains(validOperators, word) && slices.Contains(validOperators, words[idx-1]) {
// 			return 0, false
// 		} else if word == "cubed" {
// 			return 0, false
// 		} else if slices.Contains(validOperators, word) && len(values)-len(operators) == 1 {
// 			operators = append(operators, MapOperator(word))
// 		} else {
// 			num, err := strconv.Atoi(word)
// 			if err == nil {
// 				values = append(values, num)
// 			}
// 		}
// 	}

// 	if len(values) == 1 && len(operators) == 0 {
// 		return values[0], true
// 	} else if len(values) > 1 && len(operators) > 0 && len(values)-len(operators) == 1 {
// 		var ans int
// 		for idx, opr := range operators {
// 			if idx == 0 {
// 				ans = Calculator(values[idx], values[idx+1], opr)
// 			} else {
// 				ans = Calculator(ans, values[idx+1], opr)
// 			}
// 		}
// 		return ans, true
// 	} else {
// 		return 0, false
// 	}
// }

func main() {
	fmt.Println(Answer("What is 5 plus 13?"))
	fmt.Println(Answer("What is 5 plus 13 plus 6?"))
	fmt.Println(Answer("What is 5 plus 13 minus 6?"))
	fmt.Println(Answer("What is 6 plus 12 plus 6 divided by 4?"))
	fmt.Println(Answer("What is 5 cubed?"))
	fmt.Println(Answer("What is 6 plus plus 3?"))
	fmt.Println(Answer("What is plus 6 3?"))
	fmt.Println(Answer("What is -3 multiplied by 25?"))

}
