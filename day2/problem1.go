package day2

import "fmt"

func Problem1() {
	// Could bring the space complexity to O(1) by evaluating each round as it comes in
	games, err := Parse("./day2/input.txt")
	if err != nil {
		panic(err)
	}

	availableBlocks := Blockset{
		Blue:  14,
		Green: 13,
		Red:   12,
	}

	result := 0
	for _, game := range games {
		possible := true
		for _, round := range game.Rounds {
			if round.Blue > availableBlocks.Blue || round.Green > availableBlocks.Green || round.Red > availableBlocks.Red {
				possible = false
				break
			}
		}

		if possible {
			result += game.Id
			//fmt.Println(game.Id)
		}
	}
	fmt.Println(result)
}