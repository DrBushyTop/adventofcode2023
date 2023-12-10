package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func NewInput(folderPath string) *Input {
	i := Input{}
	i.Seeds = ParseSeeds(fmt.Sprintf("%s/seeds.txt", folderPath))
	i.SeedToSoil = ParseData(fmt.Sprintf("%s/seed-to-soil.txt", folderPath))
	i.SoilToFertilizer = ParseData(fmt.Sprintf("%s/soil-to-fertilizer.txt", folderPath))
	i.FertilizerToWater = ParseData(fmt.Sprintf("%s/fertilizer-to-water.txt", folderPath))
	i.WaterToLight = ParseData(fmt.Sprintf("%s/water-to-light.txt", folderPath))
	i.LightToTemperature = ParseData(fmt.Sprintf("%s/light-to-temperature.txt", folderPath))
	i.TemperatureToHumidity = ParseData(fmt.Sprintf("%s/temperature-to-humidity.txt", folderPath))
	i.HumidityToLocation = ParseData(fmt.Sprintf("%s/humidity-to-location.txt", folderPath))

	return &i
}

func ParseSeeds(path string) []int {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var res []int

	nums := strings.Split(string(f), " ")
	for _, num := range nums {
		parsed, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			panic(err)
		}
		res = append(res, parsed)
	}

	return res
}

func BruteForceParseSeeds(path string) [][]int {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var res [][]int

	nums := strings.Split(string(f), " ")
	for i, num := range nums {
		parsed, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			res = append(res, []int{parsed})
		} else {
			previous := res[len(res)-1]
			previous = append(previous, parsed)
			res[len(res)-1] = previous
		}
	}

	return res
}

func ParseData(path string) []Data {
	res := make([]Data, 0)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		destinationStart, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		sourceStart, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		res = append(res, Data{
			SourceStart:      sourceStart,
			DestinationStart: destinationStart,
			Range:            r,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].SourceStart < res[j].SourceStart
	})

	return res
}