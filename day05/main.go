package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var cards = make(map[int]int)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sections := getSections(fileScanner)

	seeds := mapToInt(strings.Fields(string(sections[0][0][7:])))

	seedToSoil := getListOfIntsFromSection(sections[1])
	soilToFertilizer := getListOfIntsFromSection(sections[2])
	fertilizerToWater := getListOfIntsFromSection(sections[3])
	waterToLight := getListOfIntsFromSection(sections[4])
	lightToTemperature := getListOfIntsFromSection(sections[5])
	temperatureToHumidity := getListOfIntsFromSection(sections[6])
	humidityToLocation := getListOfIntsFromSection(sections[7])

	seedMaps := make([][][]int, 7)
	seedMaps[0] = seedToSoil
	seedMaps[1] = soilToFertilizer
	seedMaps[2] = fertilizerToWater
	seedMaps[3] = waterToLight
	seedMaps[4] = lightToTemperature
	seedMaps[5] = temperatureToHumidity
	seedMaps[6] = humidityToLocation

	minLocation := math.MaxInt
	seedsLength := len(seeds)
	for i := 0; i < seedsLength; i += 2 {
		s := seeds[i]
		sRange := seeds[i] + seeds[i+1]
		for j := s; j < sRange; j++ {
			location := findLocation(j, seedMaps)
			if location < minLocation {
				minLocation = location
			}
		}
	}

	fmt.Printf("Minimum Location: %d\n", int(minLocation))
}

func getSections(fileScanner *bufio.Scanner) [][]string {
	var sections [][]string
	var section []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			sections = append(sections, section)
			section = nil
		} else {
			section = append(section, line)
		}
	}

	if section != nil {
		sections = append(sections, section)
	}
	section = nil
	return sections
}

func getListOfIntsFromSection(section []string) [][]int {
	numbers := make([][]int, len(section)-1)
	for i, s := range section {
		// ignore the title
		if i == 0 {
			continue
		}
		numbers[i-1] = mapToInt(strings.Fields(s))
	}
	return numbers
}

func mapToInt(slice []string) []int {
	numbers := make([]int, len(slice))
	i := 0
	for _, s := range slice {
		number, _ := strconv.Atoi(s)
		numbers[i] = number
		i++
	}
	return numbers
}

func findLocation(s int, seedMaps [][][]int) int {
	for _, seedMap := range seedMaps {
		for _, row := range seedMap {
			if s >= row[1] && s < row[1]+row[2] {
				s = row[0] + (s - row[1])
				break
			}
		}
	}

	return s
}
