package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) ([]map[int]bool, [][]int, error) {
	var winningNumbers []map[int]bool
	var myNumbers [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		numbersPart := strings.Split(scanner.Text(), ": ")[1]
		win := strings.Split(numbersPart, " | ")[0]
		my := strings.Split(numbersPart, " | ")[1]

		winningNumbers = append(winningNumbers, make(map[int]bool))
		myNumbers = append(myNumbers, make([]int, 0))

		for _, num := range strings.Split(win, " ") {
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}
			winningNumbers[i][n] = true
		}

		for _, num := range strings.Split(my, " ") {
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, err
			}
			myNumbers[i] = append(myNumbers[i], n)
		}
		i++
	}

	return winningNumbers, myNumbers, nil
}