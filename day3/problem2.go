package day3

import (
	"fmt"
	"strconv"
)

func Problem2() {
	// Could optimize the space complexity of this part a bit by keeping 3 lines at a time. Depending on if the schematics
	// are always the same size, this would bring the complexity to O(1) or O(N) where N is the number of characters on a line
	schematic, err := ParseSchematic("./day3/input.txt")
	if err != nil {
		panic(err)
	}

	gears := Traverse2(schematic)

	result := 0

	for _, numbers := range gears {
		// Only consider gears with exactly two adjacent numbers
		if len(numbers) != 2 {
			continue
		}
		result += numbers[0] * numbers[1]
	}

	fmt.Println(result)
}

type Coordinate struct {
	row, col int
}

func Traverse2(schematic [][]rune) map[Coordinate][]int {
	gears := make(map[Coordinate][]int)
	for row := 0; row < len(schematic); row++ {
		col := 0
		currentNumStr := ""
		gearLoc := Coordinate{-1, -1}
		for col < len(schematic[row]) {
			// Find next Number start
			for col < len(schematic[row]) && (schematic[row][col] < '0' || schematic[row][col] > '9') {
				col++
			}
			// We're at the first number
			// Continue until the end of the number or the end of the line
			for col < len(schematic[row]) && (schematic[row][col] >= '0' && schematic[row][col] <= '9') {
				currentNumStr += string(schematic[row][col])
				if gearLoc.row == -1 && gearLoc.col == -1 {
					gearLoc = IsPartOfGearRatio(schematic, row, col, gears)
				}
				col++
			}
			// If we're next to a gear, we need to add the number to the items adjacent to the gear
			if gearLoc.row != -1 && gearLoc.col != -1 {
				num, err := strconv.Atoi(currentNumStr)
				if err != nil {
					panic(err)
				}
				gears[gearLoc] = append(gears[gearLoc], num)
			}
			// Reset the number string and gear location
			currentNumStr = ""
			gearLoc = Coordinate{-1, -1}
			col++
		}
	}

	return gears

}

// Duplicate for readability
var Directions2 = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func IsPartOfGearRatio(schematic [][]rune, row, col int, gears map[Coordinate][]int) Coordinate {

	for _, direction := range Directions2 {
		newRow := row + direction[0]
		newCol := col + direction[1]

		if newRow < 0 || newRow >= len(schematic) {
			continue
		}

		if newCol < 0 || newCol >= len(schematic[row]) {
			continue
		}

		if schematic[newRow][newCol] == '*' {
			// Found a gear
			gLoc := Coordinate{newRow, newCol}
			if _, ok := gears[gLoc]; !ok {
				gears[gLoc] = make([]int, 0)
			}
			return gLoc
		}
	}

	return Coordinate{-1, -1}
}