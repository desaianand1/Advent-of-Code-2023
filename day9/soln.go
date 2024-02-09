// AoC Template Go file
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

func parseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error parsing number: %v\n", err)
		os.Exit(1)
	}
	return num
}

func parseSequences(lines []string) [][]int {
	sequences := [][]int{}
	for _, line := range lines {
		tokens := strings.Fields(line)
		sequence := []int{}
		for _, token := range tokens {
			sequence = append(sequence, parseInt(token))
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}

func predictExtrapolations(sequences [][]int, isPartOne bool) []int {
	extrapolations := make([]int, len(sequences[0]))
	for _, sequence := range sequences {
		extrapolations = append(extrapolations, sumSequenceDifferences(sequence, getSequenceEndValue(sequence, isPartOne), isPartOne))
	}
	return extrapolations
}

func sumSequenceDifferences(sequence []int, extrapolation int, isPartOne bool) int {
	fmt.Printf("seq: %v | sum: %d\n",sequence,extrapolation)
	sequenceSum := sum(sequence)
	if sequenceSum == 0 {
		return extrapolation
	}
	sequenceDiff := []int{}
	for i := 0; i < len(sequence)-1; i++ {
		diff := sequence[i+1] - sequence[i]
		sequenceDiff = append(sequenceDiff, diff)
	}
	return sumSequenceDifferences(sequenceDiff, extrapolation+getSequenceEndValue(sequenceDiff, isPartOne), isPartOne)
}

func getSequenceEndValue(sequence []int, shouldGetLastVal bool) int {
	if shouldGetLastVal {
		return sequence[len(sequence)-1]
	} else {
		return sequence[0]
	}
}

func sum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}
func runP1(lines []string) int {
	sequences := parseSequences(lines)
	extrapolations := predictExtrapolations(sequences, true)
	return sum(extrapolations)
}

func runP2(lines []string) int {
	sequences := parseSequences(lines)
	extrapolations := predictExtrapolations(sequences, false)
	return sum(extrapolations)
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
