package day2

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Blockset represents a single round of shown blocks
type Blockset struct {
	Blue, Green, Red int
}

// Game consists of rounds of block showings
type Game struct {
	Id     int
	Rounds []Blockset
}

func Parse(filePath string) ([]Game, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []Game{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games := make([]Game, 0)

	for scanner.Scan() {
		g, err := ParseRow(scanner.Text())
		if err != nil {
			return []Game{}, err
		}

		games = append(games, g)
	}

	return games, nil
}

func ParseRow(row string) (Game, error) {
	var game Game
	parts := strings.Split(row, ": ")
	if len(parts) != 2 {
		return game, errors.New("invalid row format")
	}

	// Parse game ID
	idPart := strings.TrimSpace(parts[0])
	idStr := strings.TrimPrefix(idPart, "Game ")
	gameId, err := strconv.Atoi(idStr)
	if err != nil {
		return game, errors.New("invalid game ID")
	}
	game.Id = gameId

	// Parse rounds
	roundsPart := parts[1]
	roundsStr := strings.Split(roundsPart, "; ")
	for _, roundStr := range roundsStr {
		var blockset Blockset
		blocks := strings.Split(roundStr, ", ")
		for _, block := range blocks {
			parts := strings.Split(block, " ")
			if len(parts) != 2 {
				return game, errors.New("invalid block format")
			}
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return game, errors.New("invalid block count")
			}
			switch parts[1] {
			case "blue":
				blockset.Blue += count
			case "green":
				blockset.Green += count
			case "red":
				blockset.Red += count
			default:
				return game, errors.New("invalid block color")
			}
		}
		game.Rounds = append(game.Rounds, blockset)
	}

	return game, nil
}