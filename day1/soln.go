package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
)

func extractDigitsP1(line string) string {
	var first string
	lastIdx := -1

	for i, el := range line {
		if el >= '0' && el <= '9' {
			if first == "" {
				first = string(el)
			} else {
				lastIdx = i
			}
		}
	}

	if first == "" {
		return ""
	}

	if lastIdx == -1 {
		return first + first
	}

	return first + string(line[lastIdx])
}

func extractDigitsP2(line string) string {
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	pattern := `(one|two|three|four|five|six|seven|eight|nine|[1-9])`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return ""
	}

	first := matches[0][1]
	second := matches[len(matches)-1][1]

	if val, ok := numMap[first]; ok {
		first = strconv.Itoa(val)
	}

	if val, ok := numMap[second]; ok {
		second = strconv.Itoa(val)
	}

	return first + second
}

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

func runP1(lines []string) {
	sum := 0
	for _, line := range lines {
		digits := extractDigitsP1(line)
		if digits != "" {
			digitVal, err := strconv.Atoi(digits)
			if err == nil {
				sum += digitVal
			}
		}
	}
	fmt.Printf("part 1: %d\n", sum)
}

func runP2(lines []string) {
	sum := 0
	for _, line := range lines {
		digits := extractDigitsP2(line)
		if digits != "" {
			digitVal, err := strconv.Atoi(digits)
			if err == nil {
				sum += digitVal
			}
		}
	}
	fmt.Printf("part 2: %d\n", sum)
}

func main() {

	lines := parseArgs()
	runP1(lines)
	runP2(lines)
}
