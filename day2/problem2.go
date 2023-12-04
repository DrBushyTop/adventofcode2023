package day2

import "fmt"

func Problem2() {
	// Could bring the space complexity to O(1) by evaluating each round as it comes in
	games, err := Parse("./day2/input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, game := range games {

		currentMax := Blockset{
			Blue:  0,
			Green: 0,
			Red:   0,
		}

		for _, round := range game.Rounds {
			currentMax.Blue = max(currentMax.Blue, round.Blue)
			currentMax.Green = max(currentMax.Green, round.Green)
			currentMax.Red = max(currentMax.Red, round.Red)
		}

		result += currentMax.Blue * currentMax.Green * currentMax.Red
	}
	fmt.Println(result)
}