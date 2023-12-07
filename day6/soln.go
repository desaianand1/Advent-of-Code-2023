package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

func getTimesAndDistances(lines []string) (times, distances []string) {
	times = strings.Fields(lines[0])[1:]
	distances = strings.Fields(lines[1])[1:]
	return times, distances
}

func makeFloat(val string) float64 {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		fmt.Printf(" %s is not an integer!", val)
		os.Exit(1)
	}
	return float64(intVal)
}

func getPossibleWays(totalTime, targetDistance float64) int {

	/** After trying linear and binary search, it seems that solving the actual equation makes the most sense
	 	h = hold time (ms)
		t = total time (ms)
		d = distance to beat (mm)
		therefore,
			h*(t - h) > d
		->  ht - h^2 - d > 0
		since we only have discrete intervals,
		->	ht - h^2 - d >= 1
		->	ht - h^2 - d - 1 >= 0
		tidied up,
		>	-h^2 + ht - d - 1 >= 0
		so, coefficients for quadratic formula ah^2 + bh + c = 0
		> a = -1 ; b = t ; c = -d-1
		to solve, h = (-b+- sqrt(b^2 - 4ac))/2a
		> h = (-t +- sqrt(t^2-4d-1))/-2
		solve for h,
		> root1: -t + sqrt(t^2-4d-1) / -2
		> root2: -t - sqrt(t^2-4d-1) / -2
	**/
	root1 := (-totalTime + math.Sqrt(math.Pow(totalTime, 2)-(4*targetDistance)-1)) / -2.0
	root2 := (-totalTime - math.Sqrt(math.Pow(totalTime, 2)-(4*targetDistance)-1)) / -2.0
	var highVal float64
	var lowVal float64

	if root1 > root2 {
		highVal = root1
		lowVal = root2
	} else {
		highVal = root2
		lowVal = root1
	}

	return int(math.Ceil(highVal) - math.Floor(lowVal) - 1)
}

func runP1(lines []string) int {
	times, distances := getTimesAndDistances(lines)
	result := 1
	for i, timeStr := range times {
		time, distance := makeFloat(timeStr), makeFloat(distances[i])
		ways := getPossibleWays(time, distance)
		result *= ways
	}
	return result
}

func runP2(lines []string) int {
	times, distances := getTimesAndDistances(lines)
	time, distance := makeFloat(strings.Join(times, "")), makeFloat(strings.Join(distances, ""))
	return getPossibleWays(time, distance)
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
