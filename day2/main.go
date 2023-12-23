package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxRed   = 12
	MaxBlue  = 14
	MaxGreen = 13
)

func main() {
	if len(os.Args) == 1 {
		panic("missing input file")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	scanner := bufio.NewScanner(file)
	sumID := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := processLine(line)
		if !game.Invalid {
			sumID += game.ID
		}
	}

	fmt.Printf("sumID: %+v\n", sumID)
	//	result was 2162
}

func processLine(line string) Game {
	s1 := strings.Split(line, ":")
	fmt.Println("*** ", s1[0])
	idStr := strings.Split(s1[0], " ")[1]
	id, _ := strconv.Atoi(idStr)
	game := Game{ID: id}
	records := strings.Split(s1[1], ";")
	for _, subset := range records {
		cubes := strings.Split(subset, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			cube = strings.ToLower(cube)
			//fmt.Printf("'%v'\n", cube)
			scores := strings.Split(cube, " ")
			v, _ := strconv.Atoi(scores[0])
			switch scores[1] {
			case "blue":
				if v > MaxBlue {
					game.Invalid = true
				}
			case "green":
				if v > MaxGreen {
					game.Invalid = true
				}
			case "red":
				if v > MaxRed {
					game.Invalid = true
				}
			}
			//fmt.Printf("%+v\n", game)
		}
	}
	return game
}

type Game struct {
	ID      int
	Invalid bool
}
