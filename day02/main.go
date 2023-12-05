package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var gameIdSum = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Scan for game id.
		gameId := getGameId(line)
		maxBlue, maxRed, maxGreen := getBlockColoursMaxCount(line)

		fmt.Printf("Game ID: %d, Max Blue: %d, Max Red: %d, Max Green: %d\n", gameId, maxBlue, maxRed, maxGreen)
		gameIdSum += maxBlue * maxRed * maxGreen
	}

	fmt.Printf("Total Sum: %d\n", gameIdSum)
}

func getGameId(text string) int {
	// start the index where the game id number begins.
	var i = 4
	var idSlice = []int{}
	for true {
		if rune(text[i]) == ':' {
			break
		} else if unicode.IsNumber(rune(text[i])) {
			idDigit, _ := strconv.Atoi(string(text[i]))
			idSlice = append(idSlice, idDigit)
		}
		i++
	}

	return combineDigits(idSlice)
}

func getBlockColoursMaxCount(text string) (blue int, red int, green int) {
	// Start index after Game ID, e.g. "Game 12: "
	var i = strings.Index(text, ":") + 2
	for true {
		count, colour, nextIndex := getBlockDetails(text, i)
		i = nextIndex
		switch colour {
		case "blue":
			blue = int(math.Max(float64(count), float64(blue)))
		case "red":
			red = int(math.Max(float64(count), float64(red)))
		case "green":
			green = int(math.Max(float64(count), float64(green)))
		}
		if i > len(text) {
			return
		}
	}
	return
}

// Get the next count and colour of the block in the given string starting
// from the given index.
// Also returns the next index for the next iteration of this function.
func getBlockDetails(text string, i int) (count int, colour string, nextIndex int) {
	var countSlice = []int{}
	var colourStartIndex = 0
	for true {
		if i >= len(text) || rune(text[i]) == ',' || rune(text[i]) == ';' {
			colour = text[colourStartIndex:i]
			nextIndex = i + 2
			return
		} else if rune(text[i]) == ' ' {
			count = combineDigits(countSlice)
			colourStartIndex = i + 1
		} else if unicode.IsNumber(rune(text[i])) {
			digit, _ := strconv.Atoi(string(text[i]))
			countSlice = append(countSlice, digit)
		}
		i++
	}
	return
}

// Join digits into a single number.
// e.g. [5, 3] -> 53.
func combineDigits(slice []int) int {
	var digitSB strings.Builder
	for _, d := range slice {
		digitSB.WriteString(fmt.Sprint(d))
	}
	number, _ := strconv.Atoi(digitSB.String())
	return number
}
