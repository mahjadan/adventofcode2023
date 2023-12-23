package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var digitWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var digitWordsRex = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)

func main() {
	if len(os.Args) == 1 {
		panic("missing input file")
	}

	// 1st way
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)
		fmt.Println("len, line: ", len(line), line)
		lineDigits := calculate(line)
		sum += lineDigits
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	slog.Info("SUCCESS", "total", sum)
	// result was 54414

	// 2nd way
	file2, err := os.Open(os.Args[1])
	if err != nil {
		log.Println("Open file error:", err)
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)

	total := 0
	for scanner2.Scan() {
		line := scanner2.Text()
		total += firstDigit(line)*10 + lastNumber(line)
	}

	fmt.Println("RESULT 2: ", total)
	//	result was 54431
}

// the idea is to save the index and the value of each char that is number, and also save the index of each digit-word,
// then get the keys
// of this map and get the min and max (first, last) and convert them to int.
func calculate(line string) int {
	indexValueMap := make(map[int]string)
	var digitsStr []string
	for inx, char := range line {
		if unicode.IsDigit(char) {
			digitsStr = append(digitsStr, strconv.Itoa(int(char-'0')))
			fmt.Println("index: ", inx)
			indexValueMap[inx] = string(char)
		}
	}
	findDigitWords(line, indexValueMap)

	// get the keys of the map to find the first and last digits
	var indexes []int
	for key := range indexValueMap {
		indexes = append(indexes, key)
	}
	firstIndex := slices.Min(indexes)
	lastIndex := slices.Max(indexes)
	resultStr := strings.Join([]string{indexValueMap[firstIndex], indexValueMap[lastIndex]}, "")

	result, _ := strconv.Atoi(resultStr)
	fmt.Println("result: ", result)
	return result
}

func findDigitWords(line string, indexValueMap map[int]string) {
	//add the matchValue and the index of first char of the word
	matchesIndices := digitWordsRex.FindAllStringIndex(line, -1)
	for _, matchIndex := range matchesIndices {
		start, end := matchIndex[0], matchIndex[1]
		match := line[start:end]
		indexValueMap[start] = digitMap[match]
	}
	fmt.Println("indexes: ", indexValueMap)
}

func firstDigit(line string) int {
	tempStr := ""

	for _, char := range line {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}

		tempStr += string(char)

		for i, digit := range digitWords {
			if strings.HasSuffix(tempStr, digit) {
				return i + 1
			}
		}
	}
	return 0
}

func lastNumber(s string) int {
	tempStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}

		tempStr = string(s[i]) + tempStr

		for j, digit := range digitWords {
			if strings.HasPrefix(tempStr, digit) {
				return j + 1
			}
		}
	}
	return 0
}
