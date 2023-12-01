package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var sum = 0
	for fileScanner.Scan() {
		var digits []rune
		line := fileScanner.Text()
		for _, r := range line {
			if unicode.IsDigit(r) {
				digits = append(digits, r)
			}
		}
		number, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		sum += number
	}
	fmt.Printf("Sum: %d", sum)
	file.Close()
}
