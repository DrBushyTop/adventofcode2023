package day7

import (
	"fmt"
	"sort"
	"strings"
)

var values2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func Problem2() {
	// Parse hands and winnings
	Hands := ParseHands("day7/input.txt")

	// Check hands for type
	var HandsByType [7][]Hand
	for _, hand := range Hands {
		HandsByType[CheckHandTypeWithJokers(hand.Cards)] = append(HandsByType[CheckHandTypeWithJokers(hand.Cards)], hand)
	}

	// Sort hands for each type to be ranked
	for i := range HandsByType {
		currIndex := i
		sort.Slice(HandsByType[currIndex], func(i, j int) bool {
			return compareHandsInRank(HandsByType[currIndex][i].Cards, HandsByType[currIndex][j].Cards, values2) > 0
		})
	}

	rankedHands := make([]Hand, 0)
	for i := len(HandsByType) - 1; i >= 0; i-- {
		rankedHands = append(rankedHands, HandsByType[i]...)
	}

	// Calculate winnings
	maxRank := len(Hands)
	winnings := 0
	for i, j := 0, maxRank; j > 0; i, j = i+1, j-1 {
		winnings += rankedHands[i].Winnings * j
	}

	fmt.Println(winnings)
}

// CheckHandTypeWithJokers checks a hand of cards and returns the highest value type found in the hand, with jokers
func CheckHandTypeWithJokers(hand string) int {
	highestDuplicates := 0

	var pairs []string

	// Map number of cards of each value
	cards := strings.Split(hand, "")
	cardMap := make(map[string]int)
	for _, card := range cards {
		cardMap[card]++
		highestDuplicates = max(highestDuplicates, cardMap[card])
		if cardMap[card] == 2 {
			pairs = append(pairs, card)
		}
	}
	numberOfJokers := cardMap["J"]

	if highestDuplicates == 5 || (highestDuplicates+numberOfJokers >= 5 && len(cardMap) <= 2) {
		return typeValues["Five of a Kind"]
	}
	if highestDuplicates == 4 || (highestDuplicates+numberOfJokers >= 4 && len(pairs) == 2) || (highestDuplicates+numberOfJokers >= 4 && numberOfJokers < 2) || (highestDuplicates+numberOfJokers >= 4 && len(cardMap) <= 3) {
		return typeValues["Four of a Kind"]
	}
	if len(pairs) == 2 && highestDuplicates == 3 || len(pairs) == 2 && highestDuplicates+numberOfJokers >= 3 {
		return typeValues["Full House"]
	}
	if highestDuplicates == 3 || highestDuplicates+numberOfJokers >= 3 {
		return typeValues["Three of a Kind"]
	}
	if len(pairs) == 2 || len(pairs) == 1 && numberOfJokers >= 1 {
		return typeValues["Two Pairs"]
	}
	if len(pairs) == 1 || len(pairs) == 0 && numberOfJokers >= 1 {
		return typeValues["One Pair"]
	}

	return typeValues["High Card"]
}