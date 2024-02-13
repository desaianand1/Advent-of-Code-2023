// AoC Template Go file
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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
	'S': Starter,
}

type PipePoint struct {
	pipe Pipe
	i    int
	j    int
}

type PipeLoop = []PipePoint

func findPipeLoop(grid []string) PipeLoop {
	var pipeLoop = PipeLoop{}
	var starterPipe PipePoint
	for i, row := range grid {
		for j, token := range row {
			direction, isDirection := pipeDirectionMap[token]
			if isDirection && direction == Starter {
				starterPipe = PipePoint{i: i, j: j, pipe: Starter}
				break
			}
		}
	}
	visitedPipes := make(map[PipePoint]bool)
	crawlPipeLoop(grid, starterPipe, &pipeLoop, &visitedPipes)
	return pipeLoop
}

func crawlPipeLoop(grid []string, point PipePoint, pipeLoop *PipeLoop, visitedPipes *map[PipePoint]bool) {
	fmt.Printf("crawling at (%d,%d): %v |\n", point.i, point.j, point.pipe)
	_, hasVisited := (*visitedPipes)[point]
	fmt.Printf("visited? (%d,%d): %v = %v\n", point.i, point.j, point.pipe, hasVisited)
	if hasVisited {
		return
	}
	(*visitedPipes)[point] = true
	*pipeLoop = append(*pipeLoop, point)
	corners := checkFourCorners(grid, point)
	fmt.Printf("corners found %v\n", corners)
	if len(corners) == 0 {
		return
	}
	for _, corner := range corners {
		if corner.pipe == Starter {
			return
		}
		crawlPipeLoop(grid, corner, pipeLoop, visitedPipes)
	}
}

func arePipesConnected(this PipePoint, other PipePoint) bool {
	// overlapping pipes not allowed
	if this.i == other.i && this.j == other.j {
		return false
	}
	// pipes are not-adjacent either vertically or horizontally
	if math.Abs(float64(this.i-other.i))-1 > 1e-9 || math.Abs(float64(this.j-other.j))-1 > 1e-9 {
		return false
	}
	// pipes are diagonally adjacent, i.e. not adjacent
	if (this.i == other.i || this.j != other.j) && (this.i != other.i || this.j == other.j) {
		return false
	}
	switch this.pipe {
	case Vertical:
		if this.i > other.i {
			return other.pipe == TopLeft || other.pipe == TopRight || other.pipe == Vertical
		} else {
			return other.pipe == BottomLeft || other.pipe == BottomRight || other.pipe == Vertical
		}
	case Horizontal:
		if this.j > other.j {
			return other.pipe == BottomLeft || other.pipe == TopLeft || other.pipe == Horizontal
		} else {
			return other.pipe == BottomRight || other.pipe == TopRight || other.pipe == Horizontal
		}
	case TopLeft:
		if this.i > other.i || this.j > other.j {
			return false
		}
		if this.j == other.j {
			return other.pipe == Vertical || other.pipe == BottomLeft || other.pipe == BottomRight
		}
		return other.pipe == Horizontal || other.pipe == BottomRight || other.pipe == TopRight

	case TopRight:
		if this.i > other.i || this.j < other.j {
			return false
		}
		if this.j == other.j {
			return other.pipe == Vertical || other.pipe == BottomLeft || other.pipe == BottomRight
		}
		return other.pipe == Horizontal || other.pipe == BottomLeft || other.pipe == TopLeft
	case BottomLeft:
		if this.i < other.i || this.j > other.j {
			return false
		}
		if this.j == other.j {
			return other.pipe == Vertical || other.pipe == TopLeft || other.pipe == TopRight
		}
		return other.pipe == Horizontal || other.pipe == BottomRight || other.pipe == TopRight
	case BottomRight:
		if this.i < other.i || this.j < other.j {
			return false
		}
		if this.j == other.j {
			return other.pipe == Vertical || other.pipe == TopLeft || other.pipe == TopRight
		}
		return other.pipe == Horizontal || other.pipe == BottomLeft || other.pipe == TopLeft
	case Starter:
		if this.i > other.i {
			return other.pipe == Vertical || other.pipe == TopLeft || other.pipe == TopRight
		} else if this.i < other.i {
			return other.pipe == Vertical || other.pipe == BottomLeft || other.pipe == BottomRight
		} else {
			// i's are equal, j's different
			if this.j > other.j {
				return other.pipe == Horizontal || other.pipe == BottomLeft || other.pipe == TopLeft
			} else {
				return other.pipe == Horizontal || other.pipe == BottomRight || other.pipe == TopRight
			}
		}
	default:
		return false
	}
}

func checkFourCorners(grid []string, point PipePoint) []PipePoint {
	above, below, left, right := point.i-1, point.i+1, point.j-1, point.j+1
	foundPoints := []PipePoint{}
	if above >= 0 {
		token := rune(grid[above][point.j])
		pipe, isPipeDirection := pipeDirectionMap[token]
		adjacentPoint := PipePoint{i: above, j: point.j, pipe: pipe}
		if isPipeDirection && arePipesConnected(point, adjacentPoint) {
			fmt.Printf("are %v and %v connected? %v\n", point, adjacentPoint, arePipesConnected(point, adjacentPoint))
			foundPoints = append(foundPoints, adjacentPoint)
		}
	}

	if below < len(grid) {
		token := rune(grid[below][point.j])
		pipe, isPipeDirection := pipeDirectionMap[token]
		adjacentPoint := PipePoint{i: below, j: point.j, pipe: pipe}
		if isPipeDirection && arePipesConnected(point, adjacentPoint) {
			foundPoints = append(foundPoints, adjacentPoint)
		}
	}

	if left >= 0 {
		token := rune(grid[point.i][left])
		pipe, isPipeDirection := pipeDirectionMap[token]
		adjacentPoint := PipePoint{i: point.i, j: left, pipe: pipe}
		if isPipeDirection && arePipesConnected(point, adjacentPoint) {
			foundPoints = append(foundPoints, adjacentPoint)
		}
	}

	if right < len(grid[point.i]) {
		token := rune(grid[point.i][right])
		pipe, isPipeDirection := pipeDirectionMap[token]
		adjacentPoint := PipePoint{i: point.i, j: right, pipe: pipe}
		if isPipeDirection && arePipesConnected(point, adjacentPoint) {
			foundPoints = append(foundPoints, adjacentPoint)
		}
	}

	return foundPoints
}

func runP1(lines []string) int {
	pipeLoop := findPipeLoop(lines)
	for _, point := range pipeLoop {
		fmt.Printf("%v\n", point.pipe)
	}
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
