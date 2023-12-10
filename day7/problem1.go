package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var values = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var typeValues = map[string]int{
	"High Card":       0,
	"One Pair":        1,
	"Two Pairs":       2,
	"Three of a Kind": 3,
	"Full House":      4,
	"Four of a Kind":  5,
	"Five of a Kind":  6,
}

type Hand struct {
	Cards    string
	Winnings int
}

func ParseHands(filepath string) []Hand {
	var hands []Hand

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		winInt, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{
			Cards:    words[0],
			Winnings: winInt,
		})
	}

	return hands
}

func Problem1() {
	// Parse hands and winnings
	Hands := ParseHands("day7/input.txt")

	// Check hands for type
	var HandsByType [7][]Hand
	for _, hand := range Hands {
		HandsByType[CheckHandType(hand.Cards)] = append(HandsByType[CheckHandType(hand.Cards)], hand)
	}

	// Sort hands for each type to be ranked
	for i := range HandsByType {
		currIndex := i
		sort.Slice(HandsByType[currIndex], func(i, j int) bool {
			return compareHandsInRank(HandsByType[currIndex][i].Cards, HandsByType[currIndex][j].Cards, values) > 0
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

// CheckHandType checks a hand of cards and returns the highest value type found in the hand
func CheckHandType(hand string) int {
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
	if highestDuplicates == 5 {
		return typeValues["Five of a Kind"]
	}
	if highestDuplicates == 4 {
		return typeValues["Four of a Kind"]
	}
	if len(pairs) == 2 && highestDuplicates == 3 {
		return typeValues["Full House"]
	}
	if highestDuplicates == 3 {
		return typeValues["Three of a Kind"]
	}
	if len(pairs) == 2 {
		return typeValues["Two Pairs"]
	}
	if len(pairs) == 1 {
		return typeValues["One Pair"]
	}

	return typeValues["High Card"]
}

// compareHandsInRank sorts a slice of hands in rank order
func compareHandsInRank(hand1, hand2 string, values map[rune]int) int {
	for i := range hand1 {
		if values[rune(hand1[i])] != values[rune(hand2[i])] {
			return values[rune(hand1[i])] - values[rune(hand2[i])]
		}
	}
	return 0
}