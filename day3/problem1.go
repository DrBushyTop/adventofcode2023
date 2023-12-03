package day3

import (
	"fmt"
	"strconv"
)

func Problem1() {
	schematic, err := ParseSchematic("./day3/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(Traverse(schematic))
}

func Traverse(schematic [][]rune) int {
	var result int
	for row := 0; row < len(schematic); row++ {
		col := 0
		currentNumStr := ""
		partNumber := false
		for col < len(schematic[row]) {
			// Find next Number start
			for col < len(schematic[row]) && (schematic[row][col] < '0' || schematic[row][col] > '9') {
				col++
			}
			// We're at the first number
			// Continue until the end of the number or the end of the line
			for col < len(schematic[row]) && (schematic[row][col] >= '0' && schematic[row][col] <= '9') {
				currentNumStr += string(schematic[row][col])
				if !partNumber {
					// For each step, scan adjacent (diagonal also) for a symbol other than '.'
					// If true at any point, identify number as a part number
					partNumber = IsPartNumber(schematic, row, col)
				}
				col++
			}
			// If at the end of the number it is a part number, add the number to the result sum
			if partNumber {
				num, err := strconv.Atoi(currentNumStr)
				if err != nil {
					panic(err)
				}
				result += num
			}
			// Reset the number string and part number flag
			currentNumStr = ""
			partNumber = false
			col++
		}
	}

	return result
}

var Directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func IsPartNumber(schematic [][]rune, row, col int) bool {

	for _, direction := range Directions {
		newRow := row + direction[0]
		newCol := col + direction[1]

		if newRow < 0 || newRow >= len(schematic) {
			continue
		}

		if newCol < 0 || newCol >= len(schematic[row]) {
			continue
		}

		if schematic[newRow][newCol] != '.' && !(schematic[newRow][newCol] >= '0' && schematic[newRow][newCol] <= '9') {
			return true
		}
	}

	return false
}