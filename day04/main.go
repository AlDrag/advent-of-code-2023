package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		id := line[4:8]
		fmt.Printf("Line #: %s\n", id)
		winningNumbersString := strings.Trim(line[9:39], "")
		chosenNumbersString := strings.Trim(line[42:], "")

		winningNumbers := mapToInt(strings.Fields(winningNumbersString))
		chosenNumbers := mapToInt(strings.Fields(chosenNumbersString))

		totalScore += calculateScore(winningNumbers, chosenNumbers)
	}

	fmt.Printf("Total Score: %d\n", totalScore)
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

func calculateScore(winningNumbers []int, chosenNumbers []int) int {
	score := 0
	for _, n := range chosenNumbers {
		if slices.Contains(winningNumbers, n) {
			if score == 0 {
				score = 1
			} else {
				score += score
			}
		}
	}
	return score
}
