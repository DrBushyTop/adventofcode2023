package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Problem1() {
	file, err := os.Open("./day1/problem.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		lineRes, err := calculateNumber(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(lineRes)
		result += lineRes
	}

	fmt.Println(result)
}

func calculateNumber(line string) (int, error) {
	// Problem 1
	first, last := 0, len(line)-1
	for line[first] < '0' || line[first] > '9' {
		first++
	}

	for line[last] < '0' || line[last] > '9' {
		last--
	}

	res, err := strconv.Atoi(fmt.Sprintf("%c%c", line[first], line[last]))
	if err != nil {
		return 0, err
	}
	return res, nil
}