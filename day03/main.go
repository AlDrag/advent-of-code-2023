package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var matrixWindow = [3][3]int{
	{-1, 0, 1},
	{-1, 0, 1},
	{-1, 0, 1},
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var matrix = [3][]rune{}
	var sum = 0

	// Populate matrix with initial 2 lines of the file (and a fake -1 line)
	fileScanner.Scan()
	line := fileScanner.Text()
	for _, r := range line {
		// Populate out of range matrix edge with full stops to simplify algorithm.
		matrix[0] = append(matrix[0], '.')
		matrix[1] = append(matrix[1], r)
	}
	fileScanner.Scan()
	line = fileScanner.Text()
	for _, r := range line {
		matrix[2] = append(matrix[2], r)
	}

	validNumbers := scanValidNumbers(line, matrix)

	for _, n := range validNumbers {
		sum += n
	}

	fmt.Printf("Line: %d, %v\n", 1, validNumbers)

	var i = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Move matrix to the next line
		matrix[0] = matrix[1]
		matrix[1] = matrix[2]
		matrix[2] = []rune{}
		for _, r := range line {
			matrix[2] = append(matrix[2], r)
		}

		resetMatrixWindow()
		validNumbers = scanValidNumbers(line, matrix)

		for _, n := range validNumbers {
			sum += n
		}

		fmt.Printf("Line: %d, %v\n", i+1, validNumbers)
		i++
	}

	matrix[0] = matrix[1]
	matrix[1] = matrix[2]
	for range line {
		// Populate out of range matrix edge with full stops to simplify algorithm.
		matrix[2] = append(matrix[2], '.')
	}

	resetMatrixWindow()
	validNumbers = scanValidNumbers(line, matrix)

	for _, n := range validNumbers {
		sum += n
	}

	fmt.Printf("Line: %v\n", validNumbers)
	fmt.Printf("Sum: %d\n", sum)
}

func scanValidNumbers(text string, matrix [3][]rune) []int {
	var validNumbers = []int{}
	var i = 0
	for i < len(text) {
		var symbol = matrix[1][i]
		if unicode.IsNumber(symbol) {
			var isValidPart = false
			var j = i
			for true && j < len(matrix[1]) {
				if unicode.IsNumber(matrix[1][j]) {
					if !isValidPart && hasCollision(matrix) {
						isValidPart = true
					}
					j++
					moveMatrixWindow(j)
				} else {
					break
				}
			}
			if isValidPart {
				var number, _ = strconv.Atoi(string(matrix[1][i:j]))
				validNumbers = append(validNumbers, number)
			}
			// Move index after the whole number
			i = j + 1
		} else {
			i++
		}

		moveMatrixWindow(i)
	}
	return validNumbers
}

func moveMatrixWindow(center int) {
	for i := range matrixWindow {
		for j := range matrixWindow[i] {
			if j == 0 {
				matrixWindow[i][j] = center - 1
			} else if j == 2 {
				matrixWindow[i][j] = center + 1
			} else {
				matrixWindow[i][j] = center
			}
		}
	}
}

func resetMatrixWindow() {
	for i := range matrixWindow {
		for j := range matrixWindow[i] {
			matrixWindow[i][j] = j - 1
		}
	}
}

func hasCollision(matrix [3][]rune) bool {
	for i, row := range matrixWindow {
		for _, j := range row {
			if j < 0 || j >= len(matrix[i]) {
				continue
			}
			r := matrix[i][j]
			if r != '.' && !unicode.IsNumber(r) {
				return true
			}
		}
	}

	return false
}
