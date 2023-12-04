package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var wordToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// Sum of all digits within the input.txt file.
	var sum = 0
	var i = 1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Scan for all digits within the line.
		digits := getDigits(line)

		// Join digits into a single number.
		// e.g. [5, 3] -> 53.
		var digitSB strings.Builder
		for _, d := range digits {
			digitSB.WriteString(fmt.Sprint(d))
		}
		number, _ := strconv.Atoi(digitSB.String())

		fmt.Printf("%d: line: %v, %d\n", i, digits, number)

		sum += number
		i++
	}
	fmt.Printf("Sum: %d\n", sum)
}

// Scan for all digits within the given string.
// The string can contain digits as ints or words, e.g. qre3ret4eight
// Returns just the first and last digits.
// If there is only 1 digit, return that as 2 duplicate elements.
// e.g. [3] -> [3, 3]
func getDigits(text string) []int {
	var digits []int
	var lineLength = len(text)
	// Loop through all characters of the given string.
	for i := 0; i < lineLength; i++ {
		var r = rune(text[i])
		// If the char is a digit or word (i.e. six),
		// convert to a number and add to digits array.
		if unicode.IsDigit(r) {
			digit, _ := strconv.Atoi(string(r))
			digits = append(digits, digit)
		} else if digit := getTranslatedDigit(text, i); digit > 0 {
			digits = append(digits, digit)
		}
	}

	// Return empty slice if no digits found.
	// Return slice of duplicate digits if only 1 number is found.
	// Return first and last digit otherwise.
	if len(digits) == 0 {
		return digits
	} else if len(digits) == 1 {
		return []int{digits[0], digits[0]}
	} else {
		return []int{digits[0], digits[len(digits)-1]}
	}
}

// Loop through each word in the wordToNumber map and compare against
// given string starting the startIndex to see if it matches the word.
// Return 0 if no match is found, as there is no "zero" in the wordToNumber
// map.
func getTranslatedDigit(text string, startIndex int) (digit int) {
	for word, digit := range wordToNumber {
		var endIndex = startIndex + len(word)
		if endIndex <= len(text) && word == text[startIndex:endIndex] {
			return digit
		}
	}
	return 0
}
