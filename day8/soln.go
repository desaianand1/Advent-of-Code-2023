// AoC Template Go file
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
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

type Instruction rune
type Instructions []Instruction

const LEFT, RIGHT Instruction = Instruction('L'), Instruction('R')

type Node string

type NodePair struct {
	first  Node
	second Node
}
type Network map[Node]NodePair

func (pair NodePair) next(instruction Instruction) Node {
	fmt.Printf("get %c of pair: %v\n", instruction, pair)

	if instruction == LEFT {
		return pair.first
	}
	return pair.second
}

func parseInstructions(line string) Instructions {
	var instructions = make([]Instruction, len(line))
	for x, ins := range line {
		instructions[x] = Instruction(ins)
	}
	return instructions
}

func parseNetwork(lines []string) Network {
	var network = make(map[Node]NodePair)
	for _, line := range lines {
		parts := strings.Split(line, "=")
		node := Node(parts[0])
		values := strings.TrimFunc(parts[1], func(r rune) bool { return r == '(' || r == ')' || r == ' ' })
		nodeValues := strings.Split(values, ",")
		network[node] = NodePair{Node(strings.TrimSpace(nodeValues[0])), Node(strings.TrimSpace(nodeValues[1]))}
	}

	return network
}

func calculateStepsRequired(network Network, instructions Instructions) int {
	numInstructions := len(instructions)
	fmt.Printf("instructions: %c\n", instructions)
	fmt.Printf("network: %v\n", network)

	current, end, instrIdx, steps := Node("AAA"), Node("ZZZ"), 0, 0
	for current != end {
		if instrIdx == numInstructions {
			instrIdx = 0
			break
		}
		var instruction = instructions[instrIdx]
		pair, doesPairExist := network[current]
		if !doesPairExist {
			fmt.Printf("Could not find %s in network %v! Aborting\n", current, network)
			os.Exit(1)
		}
		next := pair.next(instruction)
		fmt.Printf("%d. %s --%c-> %s\n", steps, current, instruction, next)
		current = next
		instrIdx += 1
		steps += 1
	}
	return steps
}

func runP1(lines []string) int {
	instructions := parseInstructions(lines[0])
	// 2 onwards to skip over blank line between instructions and network
	network := parseNetwork(lines[2:])
	return calculateStepsRequired(network, instructions)
}

func runP2(lines []string) int {
	return -1
}

func main() {
	lines := parseArgs()
	fmt.Printf("part 1: %d\n", runP1(lines))
	fmt.Printf("part 2: %d\n", runP2(lines))
}
