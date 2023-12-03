package day3

import (
	"os"
	"unicode/utf8"
)

func ParseSchematic(filePath string) ([][]rune, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var schematic [][]rune
	var row []rune
	for len(bytes) > 0 {
		runeValue, width := utf8.DecodeRune(bytes)
		bytes = bytes[width:]
		if runeValue == '\n' {
			schematic = append(schematic, row)
			row = make([]rune, 0)
			continue
		}
		row = append(row, runeValue)
	}
	schematic = append(schematic, row)
	return schematic, nil
}