package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Time   int
	Record int
}

func ParseGames(filepath string) []Game {
	var res []Game
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	//Times
	words := strings.Fields(lines[0])[1:]
	var times []int
	for _, word := range words {
		val, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		times = append(times, val)
	}

	//Records
	words = strings.Fields(lines[1])[1:]
	var records []int
	for _, word := range words {
		val, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		records = append(records, val)
	}

	for i := range times {
		game := Game{
			Time:   times[i],
			Record: records[i],
		}
		res = append(res, game)
	}

	return res
}

func Problem1() {
	games := ParseGames("day6/input.txt")

	result := 1

	for _, game := range games {
		gameWins := 0
		for i := 0; i < game.Time; i++ {
			if doIWin(i, game.Time, game.Record) {
				gameWins++
			}
		}

		result *= gameWins
	}

	fmt.Println(result)
}

func doIWin(holdTime int, totalTime, record int) bool {
	if holdTime*(totalTime-holdTime) > record {
		return true
	}
	return false
}