package main

/**
learn about switch statements
*/

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
	var card1Val int
	var card2Val int
	var dealerCardVal int

	card1Val = ParseCard(card1)
	card2Val = ParseCard(card2)
	dealerCardVal = ParseCard(dealerCard)

	switch {
	case card1Val == 11 && card2Val == 11:
		return "P"
	case card1Val+card2Val == 21:
		if dealerCardVal < 10 {
			return "W"
		} else {
			return "S"
		}
	case card1Val+card2Val >= 17 && card1Val+card2Val <= 20:
		return "S"
	case card1Val+card2Val >= 12 && card1Val+card2Val <= 16:
		if dealerCardVal >= 7 {
			return "H"
		} else {
			return "S"
		}
	case card1Val+card2Val <= 11:
		return "H"
	default:
		return "W"
	}
}

func main() {
	println(FirstTurn("ace", "ace", "jack"))   // "P"
	println(FirstTurn("ace", "king", "ace"))   // "S"
	println(FirstTurn("five", "queen", "ace")) // "H"
}
