package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type Mapping struct {
	source      int
	destination int
	length      int
}

type Range struct {
	start int
	end   int
}

type Mappings []Mapping
type Almanac []Mappings

func (mapping Mapping) apply(source int) (int, bool) {
	diff := source - mapping.source
	if diff >= 0 && diff < mapping.length {
		return mapping.destination + diff, true
	}
	return -1, false
}

func (mapping Mapping) applyRange(source Range) (before *Range, mapped *Range, after *Range) {
	start, end := mapping.source, mapping.source+mapping.length-1
	if source.start < start {
		before = &Range{start: source.start, end: min(source.end, start-1)}
	}
	if source.end >= start && source.start <= end {
		newStart, _ := mapping.apply(max(start, source.start))
		newEnd, _ := mapping.apply(min(end, source.end))
		mapped = &Range{start: newStart, end: newEnd}
	}
	if source.end > end {
		after = &Range{start: max(end+1, source.start), end: source.end}
	}

	return before, mapped, after
}

func (mappings Mappings) apply(source int) int {
	for _, mapping := range mappings {
		if diff, ok := mapping.apply(source); ok {
			return diff
		}
	}
	return source
}

func (mappings Mappings) applyRange(source Range) (ranges []Range) {
	toBeMapped := []Range{source}
	for _, mapping := range mappings {
		next := []Range{}
		for _, start := range toBeMapped {
			before, mapped, after := mapping.applyRange(start)
			if before != nil {
				next = append(next, *before)
			}
			if mapped != nil {
				ranges = append(ranges, *mapped)
			}
			if after != nil {
				next = append(next, *after)
			}
		}
		toBeMapped = next
	}
	if len(toBeMapped) > 0 {
		ranges = append(ranges, toBeMapped...)
	}
	return ranges
}

func (almanac Almanac) locations(source Range) (ranges []Range) {
	ranges = []Range{source}
	for _, mappings := range almanac {
		r := []Range{}
		for _, start := range ranges {
			r = append(r, mappings.applyRange(start)...)
		}
		ranges = r
	}
	return ranges
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
	contents, err := os.ReadFile(inputFile.Name())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	stringContents := strings.TrimSpace(string(contents))
	return strings.Split(stringContents, "\n\n")
}

func parseSeeds(line string) (seeds []int) {
	for _, seed := range strings.Fields(line)[1:] {
		seedVal, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Printf("seed %v is NOT a valid integer\n", seedVal)
			os.Exit(1)
		}
		seeds = append(seeds, seedVal)
	}
	return seeds
}

func parseMapping(line string) Mapping {
	parts := strings.Fields(line)
	destination, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Destination %v is NOT a valid integer\n", destination)
		os.Exit(1)
	}
	source, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Source %v is NOT a valid integer\n", source)
		os.Exit(1)
	}
	length, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Length %v is NOT a valid integer\n", length)
		os.Exit(1)
	}
	return Mapping{source: source, destination: destination, length: length}
}

func parseAllMappings(lines []string) (almanac Almanac) {
	for _, line := range lines {
		var mappings Mappings
		for _, section := range strings.Split(line, "\n")[1:] {
			mappings = append(mappings, parseMapping(section))
		}
		almanac = append(almanac, mappings)
	}
	return almanac
}

func runP1(lines []string) int {
	seeds := parseSeeds(lines[0])
	almanac := parseAllMappings(lines[1:])
	lowest := math.MaxInt
	for _, seed := range seeds {
		val := seed
		for _, mappings := range almanac {
			val = mappings.apply(val)
		}
		if val < lowest {
			lowest = val
		}
	}
	return lowest
}

func runP2(lines []string) int {
	seeds := parseSeeds(lines[0])
	almanac := parseAllMappings(lines[1:])
	lowest := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		locs := almanac.locations(Range{start: seeds[i], end: seeds[i] + seeds[i+1] - 1})
		for _, s := range locs {
			if s.start < lowest {
				lowest = s.start
			}
		}
	}
	return lowest
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
