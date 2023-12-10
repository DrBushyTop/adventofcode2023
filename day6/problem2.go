package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseSingleGame(filepath string) Game {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	//Times
	words := strings.Fields(lines[0])[1:]
	var time int
	var tStr string
	for _, word := range words {
		tStr += word
	}
	time, err = strconv.Atoi(tStr)
	if err != nil {
		panic(err)
	}

	//Records
	words = strings.Fields(lines[1])[1:]
	var record int
	var wStr string
	for _, word := range words {
		wStr += word
	}
	record, err = strconv.Atoi(wStr)
	if err != nil {
		panic(err)
	}

	return Game{
		Time:   time,
		Record: record,
	}
}

func Problem2() {
	game := ParseSingleGame("day6/input.txt")

	result := 0

	for i := 0; i < game.Time; i++ {
		if doIWin2(i, game.Time, game.Record) {
			result++
		}
	}

	fmt.Println(result)
}

func doIWin2(holdTime int, totalTime, record int) bool {
	if holdTime*(totalTime-holdTime) > record {
		return true
	}
	return false
}