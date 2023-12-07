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

	mapset "github.com/deckarep/golang-set/v2"
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
	defer inputFile.Close()
	return fileLines
}

func createWinningHand(winningHandRaw []string) mapset.Set[string] {
	winningHandSet := mapset.NewSet[string]()
	for _, num := range winningHandRaw {
		trimmedNum := strings.TrimSpace(num)
		if trimmedNum != "" {
			winningHandSet.Add(trimmedNum)
		}
	}
	return winningHandSet
}

func createPlayingHand(playingHandRaw []string) []string {
	playingHand := make([]string, 0)
	for _, num := range playingHandRaw {
		trimmedNum := strings.TrimSpace(num)
		if trimmedNum != "" {
			playingHand = append(playingHand, trimmedNum)
		}
	}
	return playingHand
}

func calculateMatchingNumbers(playingHand []string, winningHand mapset.Set[string]) int {
	count := 0
	for _, num := range playingHand {
		if winningHand.Contains(num) {
			count++
		}
	}
	return count
}

func runP1(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		cardHands := strings.Split(parts[1], "|")
		winningHand := createWinningHand(strings.Split(cardHands[0], " "))
		playingHand := createPlayingHand(strings.Split(cardHands[1], " "))
		points := calculateMatchingNumbers(playingHand, winningHand)
		if points != 0 {
			sum += int(math.Pow(2, float64(points-1)))
		}
	}
	return sum
}

func parseCardNumber(cardNumStr string) int {
	trimmed := strings.TrimSpace(cardNumStr)
	num, err := strconv.Atoi(trimmed)
	if err != nil {
		fmt.Printf("Error parsing card number: %v\n", err)
		os.Exit(1)
	}
	return num
}

func createCardCountMap(size int) map[int]int {
	cardCountMap := make(map[int]int, size)
	for i := 1; i < size+1; i++ {
		cardCountMap[i] = 1
	}
	return cardCountMap
}

func processAllCards(cardPointMap, cardCountMap map[int]int, currentCard int) {
	points, doesExist := cardPointMap[currentCard]
	if !doesExist || points == 0 {
		return
	}
	for i := currentCard + 1; i < currentCard+points+1; i++ {
		cardCountMap[i] += 1
		processAllCards(cardPointMap, cardCountMap, i)
	}
}

func countAllCards(cardCountMap map[int]int) int {
	sum := 0
	for _, v := range cardCountMap {
		sum += v
	}
	return sum
}
func runP2(lines []string) int {
	cardPointMap := make(map[int]int)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		cardNum := parseCardNumber(strings.Split(parts[0], "Card")[1])
		cardHands := strings.Split(parts[1], "|")
		winningHand := createWinningHand(strings.Split(cardHands[0], " "))
		playingHand := createPlayingHand(strings.Split(cardHands[1], " "))
		points := calculateMatchingNumbers(playingHand, winningHand)
		_, existsInMap := cardPointMap[cardNum]
		if !existsInMap {
			cardPointMap[cardNum] = points
		}
	}
	cardCountMap := createCardCountMap(len(cardPointMap))
	for i := range cardPointMap {
		processAllCards(cardPointMap, cardCountMap, i)
	}
	return countAllCards(cardCountMap)
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %v\n", runP1(lines))
	fmt.Printf("part 2: %v\n", runP2(lines))

}
