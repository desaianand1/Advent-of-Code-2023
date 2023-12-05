package main

import (
	"bufio"
	"flag"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"unicode"
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

func findCompleteNumbers_Keyed(lines []string) map[string]string {

	idxNumMap := make(map[string]string)

	for i, line := range lines {
		var numIdxs []string
		constructedNum := ""

		for j, ch := range line {
			isNum := isDigit(ch)

			if isNum {
				numIdxs = append(numIdxs, fmt.Sprintf("%d,%d", i, j))
				constructedNum += string(ch)
			} else if constructedNum != "" {
				// add unique suffix to numbers for discernability
				keyedNum := constructedNum + fmt.Sprintf(";key%d,%d", i, j)
				for _, idx := range numIdxs {
					idxNumMap[idx] = keyedNum
				}
				numIdxs = nil
				constructedNum = ""
			}
		}

		// Check for any remaining constructed number at the end of the line
		if constructedNum != "" {
			keyedNum := constructedNum + fmt.Sprintf(";key%d,%d", i, len(line)-1)
			for _, idx := range numIdxs {
				idxNumMap[idx] = keyedNum
			}
		}
	}

	return idxNumMap
}

func isSymbol(char rune) bool {
	return !isDigit(char) && char != '.'
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func unlockCompleteNumber(keyedNum string) (int, error) {
	parts := strings.Split(keyedNum, ";key")
	if len(parts) < 2 {
		return 0, fmt.Errorf("something wrong with keyed num!: %v | keyedNum: %s", parts, keyedNum)
	}

	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("error converting string to integer: %v", err)
	}

	return num, nil
}

func deleteAllKeysForKeyedNumber(numIdxMap map[string]string, keyedNum string) {
	for key, val := range numIdxMap {
		if keyedNum == val {
			delete(numIdxMap, key)
		}
	}
}

func runP1(lines []string) int {
	numIdxMap := findCompleteNumbers_Keyed(lines)
	sum := 0

	for i, line := range lines {
		for j, ch := range line {
			if isSymbol(ch) {
				neighbors := [...]string{
					fmt.Sprintf("%d,%d", i, j-1),
					fmt.Sprintf("%d,%d", i, j+1),
					fmt.Sprintf("%d,%d", i-1, j),
					fmt.Sprintf("%d,%d", i+1, j),
					fmt.Sprintf("%d,%d", i-1, j-1),
					fmt.Sprintf("%d,%d", i-1, j+1),
					fmt.Sprintf("%d,%d", i+1, j-1),
					fmt.Sprintf("%d,%d", i+1, j+1),
				}
				keyedNumSet := mapset.NewSet[string]()

				for _, neighbor := range neighbors {
					if keyedNum, numberIsNeighbor := numIdxMap[neighbor]; numberIsNeighbor {
						keyedNumSet.Add(keyedNum)
						deleteAllKeysForKeyedNumber(numIdxMap, keyedNum)
					}
				}

				for x := range keyedNumSet.Iter() {
					actualNum, err := unlockCompleteNumber(x)
					if err != nil {
						fmt.Printf("Error unlocking complete number: %v\n", err)
						continue
					}
					sum += actualNum
				}
			}
		}
	}

	return sum
}

func runP2(lines []string) int {
	numIdxMap := findCompleteNumbers_Keyed(lines)
	sum := 0

	for i, line := range lines {
		for j, ch := range line {
			if isSymbol(ch) {
				neighbors := [...]string{
					fmt.Sprintf("%d,%d", i, j-1),
					fmt.Sprintf("%d,%d", i, j+1),
					fmt.Sprintf("%d,%d", i-1, j),
					fmt.Sprintf("%d,%d", i+1, j),
					fmt.Sprintf("%d,%d", i-1, j-1),
					fmt.Sprintf("%d,%d", i-1, j+1),
					fmt.Sprintf("%d,%d", i+1, j-1),
					fmt.Sprintf("%d,%d", i+1, j+1),
				}

				gearPair := make([]string, 2)
				count := 0

				for _, neighbor := range neighbors {
					if keyedNum, numberIsNeighbor := numIdxMap[neighbor]; numberIsNeighbor {
						if count >= 2 {
							gearPair = gearPair[:0]
							break
						}

						gearPair[count] = keyedNum
						deleteAllKeysForKeyedNumber(numIdxMap, keyedNum)
						count++
					}
				}

				if count == 2 && len(gearPair) != 0 {
					gp1, gp2 := gearPair[0], gearPair[1]
					p1, err := unlockCompleteNumber(gp1)
					if err != nil {
						fmt.Printf("Error unlocking complete number: %v\n", err)
						continue
					}

					p2, err := unlockCompleteNumber(gp2)
					if err != nil {
						fmt.Printf("Error unlocking complete number: %v\n", err)
						continue
					}
					sum += p1 * p2
				}
			}
		}
	}
	return sum
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
