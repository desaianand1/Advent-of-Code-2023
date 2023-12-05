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

// key is unique suffix added to numbers (in case a symbol's 8-direction has the same number taking up multiple 'slots')
// this way two of the same numbers appearing around same symbol are considered uniquely discernable
func findCompleteNumbers_Keyed(lines []string) map[string]string {
	var idxNumMap = make(map[string]string)
	for i, line := range lines {
		isConstructingNum := false
		constructedNum := ""
		var numIdxs [200]string
		z := 0

		for j, ch := range line {
			isNum := isDigit(ch)
			if isConstructingNum {
				if isNum {
					numIdxs[z] = fmt.Sprintf("%d,%d", i, j)
					z++
					constructedNum += string(ch)
				} else {
					// add unique suffix to numbers (in case a symbol's 8-direction has the same number taking up multiple 'slots')
					// this way two of the same numbers appearing around same symbol are considered uniquely discernable
					keyedNum := constructedNum + fmt.Sprintf(";key%d,%d", i, j)
					for _, idx := range numIdxs {
						idxNumMap[idx] = keyedNum
					}
					isConstructingNum = false
					z = 0
					numIdxs = [200]string{}
					constructedNum = ""
				}
			} else {
				if isNum {
					numIdxs[z] = fmt.Sprintf("%d,%d", i, j)
					z++
					constructedNum += string(ch)
					isConstructingNum = true
				}
			}

			if j == len(line)-1 && isConstructingNum {
				// add unique suffix to numbers (in case a symbol's 8-direction has the same number taking up multiple 'slots')
				// this way two of the same numbers appearing around same symbol are considered uniquely discernable
				keyedNum := constructedNum + fmt.Sprintf(";key%d,%d", i, j)
				for _, idx := range numIdxs {
					idxNumMap[idx] = keyedNum
				}
				isConstructingNum = false
				z = 0
				constructedNum = ""
			}

		}
	}
	return idxNumMap
}

func isSymbol(char rune) bool {
	return !(isDigit(char) || char == '.')
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func unlockCompleteNumber(keyedNum string) int {
	parts := strings.Split(keyedNum, ";key")
	if len(parts) < 2 {
		fmt.Printf("something wrong with keyed num!: %v | keyedNum: %s\n", parts, keyedNum)
	}
	num, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return num
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
					keyedNum, numberIsNeighbor := numIdxMap[neighbor]
					if numberIsNeighbor {
						keyedNumSet.Add(keyedNum)
						deleteAllKeysForKeyedNumber(numIdxMap, keyedNum)
					}
				}
				for x := range keyedNumSet.Iter() {
					actualNum := unlockCompleteNumber(x)
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
					keyedNum, numberIsNeighbor := numIdxMap[neighbor]
					if numberIsNeighbor {
						if count >= 2 {
							gearPair = gearPair[:0]
							break
						} else {
							gearPair[count] = keyedNum
							deleteAllKeysForKeyedNumber(numIdxMap, keyedNum)
							count++
						}

					}
				}
				if count == 2 && len(gearPair) != 0 {
					gp1, gp2 := gearPair[0], gearPair[1]
					p1 := unlockCompleteNumber(gp1)
					p2 := unlockCompleteNumber(gp2)
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
