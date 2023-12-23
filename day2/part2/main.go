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
	var powerSum int64
	for scanner.Scan() {
		line := scanner.Text()
		game := processLine(line)
		powerSum += game.PowerSet
		fmt.Printf("'%+v'\n", game)

	}

	fmt.Printf("powerSum: %+v\n", powerSum)
	//	result was 72513
}

func processLine(line string) Game {
	s1 := strings.Split(line, ":")
	fmt.Println("*** ", s1[0])
	idStr := strings.Split(s1[0], " ")[1]
	id, _ := strconv.Atoi(idStr)
	game := Game{ID: id}
	records := strings.Split(s1[1], ";")
	RedUsed, GreenUsed, BlueUsed := 0, 0, 0
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
				BMin := max(0, v-BlueUsed)
				game.Blue += BMin
				BlueUsed += BMin
			case "green":
				GMin := max(0, v-GreenUsed)
				game.Green += GMin
				GreenUsed += GMin
			case "red":
				RMin := max(0, v-RedUsed)
				game.Red += RMin
				RedUsed += RMin
			}
			//fmt.Printf("%+v\n", game)
		}
	}
	game.PowerSet = int64(game.Red * game.Blue * game.Green)
	return game
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Game struct {
	ID       int
	Red      int
	Green    int
	Blue     int
	PowerSet int64
}
