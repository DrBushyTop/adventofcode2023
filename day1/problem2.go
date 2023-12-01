package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Problem2() {
	// Setup trie
	setup()

	file, err := os.Open("./day1/problem.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		lineRes, err := calculateNumber2(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(lineRes)
		result += lineRes
	}

	fmt.Println(result)
}

var Numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var Trie *TrieNode

func setup() {
	Trie = NewTrieNode()
	for k := range Numbers {
		Trie.Insert(k)
		Trie.Insert(reverseString(k))
	}
}

func reverseString(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func calculateNumber2(line string) (int, error) {
	// Problem 2
	first, last := 0, len(line)-1

	var firstNum int
	var lastNum int

	var err error

	for firstNum == 0 {
		if line[first] >= '0' && line[first] <= '9' {
			firstNum, err = strconv.Atoi(fmt.Sprintf("%c", line[first]))
			if err != nil {
				return 0, err
			}
			break
		}

		// Get the string until we find a number
		temp := first
		for temp < len(line) && (line[temp] < '0' || line[temp] > '9') {
			temp++
		}
		res, err := findSubString(line[first:temp], false)
		if err != nil {
			first = temp
			continue
		}
		firstNum = res
	}

	for lastNum == 0 && last >= 0 {
		if line[last] >= '0' && line[last] <= '9' {
			lastNum, err = strconv.Atoi(fmt.Sprintf("%c", line[last]))
			if err != nil {
				return 0, err
			}
			break
		}

		// Get the string until we find a number
		temp := last
		for temp >= 0 && (line[temp] < '0' || line[temp] > '9') {
			temp--
		}
		res, err := findSubString(line[temp+1:last+1], true)
		if err != nil {
			last = temp
			continue
		}
		lastNum = res
	}

	res, err := strconv.Atoi(fmt.Sprintf("%d%d", firstNum, lastNum))
	if err != nil {
		return 0, err
	}
	return res, nil
}

func findSubString(possibleNumber string, reverse bool) (int, error) {
	if reverse {
		possibleNumber = reverseString(possibleNumber)
	}

	for i := 0; i < len(possibleNumber); i++ {
		substr := possibleNumber[i:]
		if len(substr) < 3 {
			return 0, errors.New("not enough characters")
		}

		for j := 0; j <= len(substr); j++ {
			if Trie.Search(substr[:j]) {
				if reverse {
					return Numbers[reverseString(substr[:j])], nil
				}
				return Numbers[substr[:j]], nil
			}
			continue
		}
	}

	return 0, errors.New("no number found")
}