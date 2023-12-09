package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cards = make(map[int]int)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		id, _ := strconv.Atoi(strings.Fields(line[4:8])[0])
		fmt.Printf("Line #: %d\n", id)
		winningNumbersString := strings.Trim(line[9:39], "")
		chosenNumbersString := strings.Trim(line[42:], "")

		winningNumbers := mapToInt(strings.Fields(winningNumbersString))
		chosenNumbers := mapToInt(strings.Fields(chosenNumbersString))

		if _, ok := cards[id]; ok {
			cards[id] += 1
		} else {
			cards[id] = 1
		}

		total := calculateNumbers(winningNumbers, chosenNumbers)
		fmt.Printf("Total: %d\n", total)

		for i := 1; i <= total; i++ {
			if _, ok := cards[i+id]; ok {
				cards[i+id] += cards[id]
			} else {
				cards[i+id] = cards[id]
			}
		}
	}

	totalCards := 0
	for _, t := range cards {
		totalCards += t
	}

	fmt.Printf("Total Cards: %d\n", totalCards)
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

func calculateNumbers(winningNumbers []int, chosenNumbers []int) int {
	count := 0
	for _, n := range chosenNumbers {
		if slices.Contains(winningNumbers, n) {
			count++
		}
	}

	return count
}
