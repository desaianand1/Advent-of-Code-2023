package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func parseArgs() []string {

	input := flag.String("input", "input.txt", "input file (.txt) to be read")
	flag.Parse()
	_, currentFilePath, _, _ := runtime.Caller(0)
	dirPath := path.Dir(currentFilePath)
	inputPath := path.Join(dirPath, *input)
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	inputFile.Close()
	return fileLines
}

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

type Cube map[Color]int

type Game struct {
	id   int
	sets []Cube
}

func parseGames(lines []string) []Game {
	var games []Game
	for _, line := range lines {
		sections := strings.Split(line, ":")
		idStr := strings.Split(sections[0], " ")[1]
		gameId, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		var cubeSets []Cube
		sets := strings.Split(sections[1], ";")
		for _, set := range sets {

			cubes := strings.Split(set, ",")
			cubeMap := make(Cube)
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				pair := strings.Split(cube, " ")
				color := Color(pair[1])
				count, err := strconv.Atoi(pair[0])
				if err != nil {
					panic(err)
				}
				cubeMap[color] = count
				cubeSets = append(cubeSets, cubeMap)
			}
		}
		games = append(games, Game{id: gameId, sets: cubeSets})
	}
	return games
}

func runP1(games []Game) {
	possibleCubes := Cube{Red: 12, Green: 13, Blue: 14}
	sum := 0
	for _, game := range games {
		isPossible := true
		for _, cube := range game.sets {
			for color, count := range cube {
				colorCount, colorExists := possibleCubes[color]
				if colorExists && colorCount < count {
					isPossible = false
					break
				}
			}
		}
		if isPossible {
			sum = sum + game.id
		}
	}
	fmt.Printf("part 1: %v\n", sum)
}

func runP2(games []Game) {
	sum := 0
	for _, game := range games {
		leastCubesRequired := Cube{Red: 0, Green: 0, Blue: 0}
		power := 1
		for _, cube := range game.sets {
			for color, count := range cube {
				colorCount, colorExists := leastCubesRequired[color]
				if colorExists && colorCount < count {
					leastCubesRequired[color] = count
					break
				}
			}
		}
		for _, count := range leastCubesRequired {
			power = power * count
		}
		sum = sum + power
	}
	fmt.Printf("part 2: %v\n", sum)
}

func main() {
	lines := parseArgs()
	games := parseGames(lines)
	runP1(games)
	runP2(games)
}
