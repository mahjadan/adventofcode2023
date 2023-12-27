package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) == 1 {
		panic("missing input file")
	}
	fileName := os.Args[1]

	var grid [][]rune
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("error reading file: %v", err))
	}

	grid = parseLines(string(content))

	partNumbers := processGrid(grid)
	var sum int64
	for _, number := range partNumbers {
		sum += number
	}
	fmt.Println("sum: ", sum)
	//	result gear ratios sum 79844424

}
func processGrid(grid [][]rune) []int64 {
	var gearGrid = make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		line := grid[i]
		fmt.Println(string(line))
		gearGrid[i] = getGearIndexes(line)
	}
	fmt.Println("gearGrid: ", gearGrid)
	var gearRatios []int64
	for i, gearIndex := range gearGrid {
		for _, j := range gearIndex {
			var gears []int
			fmt.Printf("%v/%v: %v\n", i+1, j, string(grid[i][j]))
			if n := checkLine(grid, i, j); len(n) != 0 {
				gears = append(gears, n...)
			}
			if n := checkLine(grid, i+1, j); len(n) != 0 {
				gears = append(gears, n...)
			}
			if n := checkLine(grid, i-1, j); len(n) != 0 {
				gears = append(gears, n...)
			}
			if len(gears) == 2 {
				fmt.Printf("gears found near line:%v, position:%v ,gearRatio: %v ,gears: %v\n", i+1, j, gears[0]*gears[1], gears)
				gearRatios = append(gearRatios, int64(gears[0]*gears[1]))
			}
		}

	}
	fmt.Printf("gearRatios: %v\n", gearRatios)
	return gearRatios
}

func checkLine(grid [][]rune, i, j int) []int {
	var numbers []int
	if i < len(grid) && i >= 0 {
		line := grid[i]
		//fmt.Println("line: ", string(line))
		numbers = findNumberInLine(j, line)
		//fmt.Printf("numbers found near j:%v ,number: %v\n", j, numbers)
		return numbers
	}
	return nil
}

func getGearIndexes(line []rune) []int {
	var specialGearIndex []int
	for i, char := range line {
		if char == '*' {
			specialGearIndex = append(specialGearIndex, i)
		}
	}
	return specialGearIndex
}

func findNumberInLine(index int, line []rune) []int {
	var numbers []int
	// check for number right side
	numRight := checkNumberRight(index, line)

	// check for number left side
	numLeft := checkNumberLeft(index, line)

	// if the index itself is a number (in case of upper or bottom line) need to join the result
	if unicode.IsDigit(line[index]) {
		s := numLeft + string(line[index]) + numRight
		result, _ := strconv.Atoi(s)
		return []int{result}
	} else {
		if numRight != "" {
			result, _ := strconv.Atoi(numRight)
			numbers = append(numbers, result)
		}
		if numLeft != "" {
			result, _ := strconv.Atoi(numLeft)
			numbers = append(numbers, result)
		}
	}
	return numbers
}

func checkNumberLeft(index int, line []rune) string {
	var currentIndex int
	var number string
	//fmt.Printf("length/index-1: %v/%v, line: %v\n", len(line), index-1, string(line))
	if index-1 >= 0 && unicode.IsDigit(line[index-1]) {
		currentIndex = index - 1
		for currentIndex >= 0 {
			if unicode.IsDigit(line[currentIndex]) {
				number = string(line[currentIndex]) + number
				currentIndex -= 1
			} else {
				break
			}
		}
	}
	return number
}
func checkNumberRight(index int, line []rune) string {
	var currentIndex int
	var number []rune
	if index+1 <= len(line) && unicode.IsDigit(line[index+1]) {
		currentIndex = index + 1
		for currentIndex < len(line) {
			if unicode.IsDigit(line[currentIndex]) {
				number = append(number, line[currentIndex])
				currentIndex += 1
			} else {
				break
			}
		}
	}
	return string(number)
}
func parseLines(content string) [][]rune {

	lines := strings.Split(strings.TrimSpace(content), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)

	}
	return grid
}
