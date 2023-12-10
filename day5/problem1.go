package day5

import (
	"fmt"
	"math"
)

type Data struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

type Input struct {
	Seeds                 []int
	SeedToSoil            []Data
	SoilToFertilizer      []Data
	FertilizerToWater     []Data
	WaterToLight          []Data
	LightToTemperature    []Data
	TemperatureToHumidity []Data
	HumidityToLocation    []Data
}

func Problem1() {
	input := NewInput("./day5/input")

	lowestLocation := math.MaxInt

	for _, seed := range input.Seeds {
		soil := getNext(seed, input.SeedToSoil)
		fertilizer := getNext(soil, input.SoilToFertilizer)
		water := getNext(fertilizer, input.FertilizerToWater)
		light := getNext(water, input.WaterToLight)
		temperature := getNext(light, input.LightToTemperature)
		humidity := getNext(temperature, input.TemperatureToHumidity)
		location := getNext(humidity, input.HumidityToLocation)

		lowestLocation = min(lowestLocation, location)
	}

	fmt.Println(lowestLocation)
}

func getNext(source int, resourceMap []Data) int {
	// Get the indexes of the lower source start range
	lowIndex := binarySearch(source, resourceMap)
	lowMax := resourceMap[lowIndex].SourceStart + resourceMap[lowIndex].Range - 1 // Get the max value of the low range

	// Check if source is between the low and lowMax
	if source >= resourceMap[lowIndex].SourceStart && source <= lowMax {
		return source + (resourceMap[lowIndex].DestinationStart - resourceMap[lowIndex].SourceStart)
	}

	// Source not in our list ranges, return it back
	return source

}

// binarySearch returns the index of the element with the closest smaller starting range to the source
func binarySearch(source int, resourceMap []Data) int {
	left, right := 0, len(resourceMap)-1
	var mid int
	closest := -1 // Initialize to -1 to handle the case when no element is found.

	for left <= right {
		mid = left + (right-left)/2

		if resourceMap[mid].SourceStart == source {
			return mid
		}

		if resourceMap[mid].SourceStart < source {
			closest = mid // Update closest as we found a smaller SourceStart.
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// If closest is not updated (remains -1), return the index of the smallest SourceStart element.
	if closest == -1 {
		return 0
	}

	// Return the closest smaller SourceStart, if available.
	return closest
}