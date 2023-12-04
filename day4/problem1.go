package day4

func Problem1() {
	winningNumbers, myNumbers, err := Parse("./day4/input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	// Check each card
	for i, nums := range myNumbers {
		cardResult := 0
		for _, num := range nums {
			if ok, _ := winningNumbers[i][num]; ok {
				if cardResult == 0 {
					cardResult++
				} else {
					cardResult = cardResult * 2
				}
			}
		}
		result = result + cardResult
	}
	println(result)
}