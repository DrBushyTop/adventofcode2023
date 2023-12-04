package day4

func Problem2() {
	winningNumbers, myNumbers, err := Parse("./day4/input.txt")
	if err != nil {
		panic(err)
	}

	tableLen := len(myNumbers)
	numCards := make([]int, tableLen)
	for i := 0; i < tableLen; i++ {
		numCards[i] = 1
	}

	// Check each card
	for i, nums := range myNumbers {
		cardResult := 0
		for _, num := range nums {
			if ok, _ := winningNumbers[i][num]; ok {
				cardResult++
			}
		}

		for j := i + 1; j < tableLen && j <= i+cardResult; j++ {
			numCards[j] = numCards[j] + numCards[i]
		}
	}

	result := 0
	for _, num := range numCards {
		result = result + num
	}
	println(result)
}