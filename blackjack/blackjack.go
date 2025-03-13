package blackjack

import "strings"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	result := 0
	switch strings.ToLower(card) {
	case "ace":
		result = 11
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	case "ten":
		result = 10
	case "jack":
		result = 10
	case "queen":
		result = 10
	case "king":
		result = 10
	}
	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1value := ParseCard(card1)
	card2value := ParseCard(card2)
	cardSum := card1value + card2value
	dealerCardvalue := ParseCard(dealerCard)

	switch {

	case cardSum <= 11:
		return "H"

	case cardSum <= 16:
		if dealerCardvalue >= 7 {
			return "H"
		} else {
			return "S"
		}

	case cardSum <= 20:
		return "S"

	case cardSum == 21:
		if dealerCardvalue < 10 {
			return "W"
		} else {
			return "S"
		}

	default:
		return "P"
	}
}
