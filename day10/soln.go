// AoC Template Go file
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
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

type Pipe string

const (
	Vertical    Pipe = "vertical"
	Horizontal  Pipe = "horizontal"
	BottomLeft  Pipe = "bottom-left"
	BottomRight Pipe = "bottom-right"
	TopRight    Pipe = "top-right"
	TopLeft     Pipe = "top-left"
	Ground      Pipe = "ground"
	Starter     Pipe = "starter"
)

var pipeDirectionMap = map[rune]Pipe{
	'|': Vertical,
	'-': Horizontal,
	'L': BottomLeft,
	'J': BottomRight,
	'7': TopRight,
	'F': TopLeft,
	'.': Ground,
	'S': Starter,
}

func findPipeLoop(lines []string) {

	for i, line := range lines {
		for j, token := range line {
			direction, doesDirectionExist := pipeDirectionMap[token]
			if doesDirectionExist && direction == Starter {
				fmt.Printf("S position: %d, %d\n", i, j)
			}
		}
	}
}

func runP1(lines []string) int {
	findPipeLoop(lines)
	return -1
}

func runP2(lines []string) int {
	return -1
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
