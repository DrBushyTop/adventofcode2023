package day5

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Turns out my previous parsing was dumb, let's try again

type InputData struct {
	// Seeds is a range of seeds with start and end values
	Seeds [][]int
	// Blocks are the blocks of data in the correct order
	Blocks [][][]int
}

func ParseToBlocks(filePath string) InputData {
	var res InputData
	b, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	content := string(b)
	blocks := strings.Split(content, "\n\n")
	for i, block := range blocks {
		if i == 0 {
			actualValues := strings.Split(block, ": ")
			// seeds
			seeds := strings.Split(actualValues[1], " ")
			for j := 0; j < len(seeds); j = j + 2 {
				start, err := strconv.Atoi(seeds[j])
				if err != nil {
					panic(err)
				}
				count, err := strconv.Atoi(seeds[j+1])
				if err != nil {
					panic(err)
				}
				res.Seeds = append(res.Seeds, []int{start, start + count - 1})
			}
			continue
		}

		// blocks
		var data [][]int
		lines := strings.Split(block, "\n")
		for _, line := range lines[1:] {
			var lineData []int
			nums := strings.Split(line, " ")
			for _, num := range nums {
				parsed, err := strconv.Atoi(strings.TrimSpace(num))
				if err != nil {
					panic(err)
				}
				lineData = append(lineData, parsed)
			}
			data = append(data, lineData)
		}

		res.Blocks = append(res.Blocks, data)
		//fmt.Println()
		//fmt.Println(block)
	}
	return res
}

func Problem2() {
	input := ParseToBlocks("./day5/input/original.txt")

	currentLayer := input.Seeds
	for _, block := range input.Blocks {
		var newLayer [][]int
		for len(currentLayer) > 0 {
			start, end := currentLayer[len(currentLayer)-1][0], currentLayer[len(currentLayer)-1][1]
			currentLayer = currentLayer[:len(currentLayer)-1]
			overlapFound := false
			for _, r := range block {
				a, b, c := r[0], r[1], r[2]

				// Get overlap
				oStart := max(start, b)
				oEnd := min(end, b+c-1)

				if oStart < oEnd {
					overlapFound = true
					newLayer = append(newLayer, []int{oStart - b + a, oEnd - b + a}) // Transform to calculate values that match next layer mapping
					if oStart > start {
						// Add the left side of the range
						currentLayer = append(currentLayer, []int{start, oStart - 1})
					}
					if oEnd < end {
						// Add the right side of the range
						currentLayer = append(currentLayer, []int{oEnd + 1, end})
					}
					break
				}
			}
			// No overlap, add the whole range
			if !overlapFound {
				newLayer = append(newLayer, []int{start, end})
			}
		}
		currentLayer = newLayer

	}

	sort.Slice(currentLayer, func(i, j int) bool {
		return currentLayer[i][0] < currentLayer[j][0]
	})

	fmt.Println(currentLayer[0][0])
}