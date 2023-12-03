package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"regexp"
	"io"
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


func main() {
	
	if len(os.Args) < 1 {
        fmt.Println("Usage : " + os.Args[0] + " file name")
        os.Exit(1)
    }

    file, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Cannot read the file")
        os.Exit(1)
    }
    // do something with the file
    fmt.Print(string(file))

	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		digits := extractDigitsP2(line)
		// fmt.Println("digit: " + digits)
		if digits != "" {
			digitVal, err := strconv.Atoi(digits)
			if err == nil {
				sum += digitVal
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("sum: %d\n", sum)
}