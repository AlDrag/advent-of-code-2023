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

	sum += scanGearRatioSum(line, matrix)

	fmt.Printf("Line: %d, %v\n", 1, sum)

	var i = 2
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
		sum += scanGearRatioSum(line, matrix)

		fmt.Printf("Line: %d, %v\n", i, sum)
		i++
	}

	matrix[0] = matrix[1]
	matrix[1] = matrix[2]
	for range line {
		// Populate out of range matrix edge with full stops to simplify algorithm.
		matrix[2] = append(matrix[2], '.')
	}

	resetMatrixWindow()
	sum += scanGearRatioSum(line, matrix)

	fmt.Printf("Line: %v\n", sum)
	fmt.Printf("Sum: %d\n", sum)
}

func scanGearRatioSum(text string, matrix [3][]rune) int {
	var gearRatioSum = 0
	var i = 0
	for i < len(text) {
		var symbol = matrix[1][i]
		if symbol == '*' {
			gearRatioSum += getGearRatio(matrix)
		}

		i++

		moveMatrixWindow(i)
	}
	return gearRatioSum
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

func getGearRatio(matrix [3][]rune) (ratio int) {
	var ratios []int
	for i, row := range matrixWindow {
		var leftIndexScanned int
		var rightIndexScanned int
		for _, j := range row {
			if j < 0 || j >= len(matrix[i]) || (j >= leftIndexScanned && j <= rightIndexScanned) {
				continue
			}
			r := matrix[i][j]
			if unicode.IsNumber(r) {
				var ratio = 0
				ratio, leftIndexScanned, rightIndexScanned = scanNumber(matrix[i], j)
				ratios = append(ratios, ratio)
			}
		}
	}

	if len(ratios) == 2 {
		ratio = ratios[0] * ratios[1]
	} else {
		ratio = 0
	}

	fmt.Printf("Gear Ratio: %d\n", ratio)

	return
}

func scanNumber(runes []rune, startingIndex int) (number int, leftIndexScanned int, rightIndexScanned int) {
	var r rune
	leftIndex := startingIndex
	rightIndex := startingIndex

	for leftIndex-1 >= 0 {
		r = runes[leftIndex-1]
		if !unicode.IsNumber(r) {
			break
		}
		leftIndex--
	}

	for rightIndex+1 < len(runes) {
		r = runes[rightIndex+1]
		if !unicode.IsNumber(r) {
			break
		}
		rightIndex++
	}

	number, _ = strconv.Atoi(string(runes[leftIndex : rightIndex+1]))
	fmt.Printf("Scanned Runes: %d\n", number)
	return number, leftIndex, rightIndex
}
