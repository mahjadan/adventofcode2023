package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	slog.Info("args: ", os.Args)
	if len(os.Args) == 1 {
		panic("missing input file")
	}

	// check for input file as argument
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read first line
		line := scanner.Text()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lineDigits := getDigits(line)
		sum += lineDigits
	}
	slog.Info("SUCCESS", "total", sum)
	//	result 55477
}
func getDigits(line string) int {
	var digitsStr []string
	for _, char := range line {
		if unicode.IsDigit(char) {
			digitsStr = append(digitsStr, string(char))
		}
	}
	first := digitsStr[:1]
	last := digitsStr[len(digitsStr)-1:]
	resultStr := strings.Join([]string{first[0], last[0]}, "")
	result, _ := strconv.Atoi(resultStr)
	return result
}
