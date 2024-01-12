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
	"sort"
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

type HandType string

const (
	FIVE_OF_A_KIND  HandType = "FIVE_OF_A_KIND"
	FOUR_OF_A_KIND  HandType = "FOUR_OF_A_KIND"
	FULL_HOUSE      HandType = "FULL_HOUSE"
	THREE_OF_A_KIND HandType = "THREE_OF_A_KIND"
	TWO_PAIR        HandType = "TWO_PAIR"
	ONE_PAIR        HandType = "ONE_PAIR"
	HIGH_CARD       HandType = "HIGH_CARD"
)

var cardRankMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardHandTypeMap = map[HandType]int{
	FIVE_OF_A_KIND:  7,
	FOUR_OF_A_KIND:  6,
	FULL_HOUSE:      5,
	THREE_OF_A_KIND: 4,
	TWO_PAIR:        3,
	ONE_PAIR:        2,
	HIGH_CARD:       1,
}

type CardHand struct {
	cards string
	bid   int
	_type HandType
}

type byHand []CardHand

func (hands byHand) Len() int {
	return len(hands)
}
func (hands byHand) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}
func (hands byHand) Less(i, j int) bool {
	var thisHand, otherHand = hands[i], hands[j]
	var comparedValue = thisHand._type.compareTo(otherHand._type)
	if comparedValue != 0 {
		return comparedValue < 0
	} else {
		thisHandRunes, otherHandRunes := []rune(thisHand.cards), []rune(otherHand.cards)
		for i, cardRune := range thisHandRunes {
			thisVal, otherVal := cardRankMap[cardRune], cardRankMap[otherHandRunes[i]]
			if thisVal != otherVal {
				return thisVal < otherVal
			}
		}
		return false
	}
}

// returns 1 if this hand type is better than the other. -1 if other hand type is better. 0 if both hand types are equal
func (thisHandType HandType) compareTo(otherHandType HandType) int {
	if thisHandType == otherHandType {
		return 0
	}
	var thisVal, otherVal int = cardHandTypeMap[thisHandType], cardHandTypeMap[otherHandType]
	if thisVal < otherVal {
		return -1
	} else {
		return 1
	}
}

func parseInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%v is NOT a valid integer\n", integer)
		os.Exit(1)
	}
	return integer
}

func createCardCountMap(cards string) map[rune]int {
	cardMap := make(map[rune]int)
	for _, card := range cards {
		_, cardExists := cardMap[card]
		if cardExists {
			cardMap[card] += 1
		} else {
			cardMap[card] = 1
		}
	}
	return cardMap
}

func determineHandType(cards string) HandType {
	cardMap := createCardCountMap(cards)
	maxCount := math.MinInt32
	var highCountCard rune
	for card, count := range cardMap {
		if count > maxCount {
			maxCount = count
			highCountCard = card
		}
	}

	const FIVE_COUNT, FOUR_COUNT, THREE_COUNT, TWO_COUNT, HIGH int = 5, 4, 3, 2, 1
	switch maxCount {
	case FIVE_COUNT:
		return FIVE_OF_A_KIND
	case FOUR_COUNT:
		return FOUR_OF_A_KIND
	case THREE_COUNT:
		// check if full house or three of a kind
		for _, count := range cardMap {
			if count == TWO_COUNT {
				return FULL_HOUSE
			}
		}
		return THREE_OF_A_KIND
	case TWO_COUNT:
		// check if two pair or just one pair
		for card, count := range cardMap {
			if card != highCountCard && count == TWO_COUNT {
				return TWO_PAIR
			}
		}
		return ONE_PAIR
	case HIGH:
		return HIGH_CARD
	default:
		return HIGH_CARD
	}
}

func determineHandTypeP2(cards string) HandType {
	cardMap := createCardCountMap(cards)
	maxCount, jCount := math.MinInt32, 0
	var highCountCard rune
	for card, count := range cardMap {
		if card == 'J' {
			jCount = count
		} else if count > maxCount {
			maxCount = count
			highCountCard = card
		}
	}

	const FIVE_COUNT, FOUR_COUNT, THREE_COUNT, TWO_COUNT, HIGH int = 5, 4, 3, 2, 1

	if maxCount == math.MinInt32 && jCount == FIVE_COUNT {
		return FIVE_OF_A_KIND
	}

	switch maxCount {
	case FIVE_COUNT:
		return FIVE_OF_A_KIND
	case FOUR_COUNT:
		switch jCount {
		case HIGH:
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	case THREE_COUNT:

		switch jCount {
		case TWO_COUNT:
			return FIVE_OF_A_KIND
		case HIGH:
			return FOUR_OF_A_KIND
		}
		// check if full house or three of a kind
		for _, count := range cardMap {
			if count == TWO_COUNT {
				return FULL_HOUSE
			}
		}
		return THREE_OF_A_KIND
	case TWO_COUNT:

		switch jCount {
		case THREE_COUNT:
			return FIVE_OF_A_KIND
		case TWO_COUNT:
			return FOUR_OF_A_KIND
		case HIGH:
			for card, count := range cardMap {
				if card != highCountCard && count == TWO_COUNT {
					return FULL_HOUSE
				}
			}
			return THREE_OF_A_KIND
		}
		// check if two pair or just one pair
		for card, count := range cardMap {
			if card != highCountCard && count == TWO_COUNT {
				return TWO_PAIR
			}
		}
		return ONE_PAIR
	case HIGH:
		switch jCount {
		case FOUR_COUNT:
			return FIVE_OF_A_KIND
		case THREE_COUNT:
			return FOUR_OF_A_KIND
		case TWO_COUNT:
			return THREE_OF_A_KIND
		case HIGH:
			return ONE_PAIR
		}
		return HIGH_CARD
	default:
		return HIGH_CARD
	}
}

func parseHands(lines []string, isPartTwo bool) []CardHand {
	hands := make([]CardHand, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		cards, bid := parts[0], parseInt(parts[1])
		var handType HandType
		if isPartTwo {
			handType = determineHandTypeP2(cards)
		} else {
			handType = determineHandType(cards)
		}
		hands[i] = CardHand{
			cards: cards,
			bid:   bid,
			_type: handType,
		}
	}
	return hands
}

func calculateTotalWinnings(rankedHands []CardHand) int {
	total := 0
	for i, hand := range rankedHands {
		total += (i + 1) * hand.bid
	}
	return total
}

func rankHands(hands []CardHand) []CardHand {
	sort.Sort(byHand(hands))
	return hands
}

func runP1(lines []string) int {
	cardRankMap['J'] = 11
	var hands = parseHands(lines, false)
	rankedHands := rankHands(hands)
	return calculateTotalWinnings(rankedHands)
}

func runP2(lines []string) int {
	cardRankMap['J'] = 1
	var hands = parseHands(lines, true)
	rankedHands := rankHands(hands)
	return calculateTotalWinnings(rankedHands)
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
