// AoC Template Go file
package main

import (
	"flag"
)

func parse_args() string {
	input := flag.String("input", "input.txt", "input file (.txt) to be read")
	flag.Parse()
	return *input
}

func main() {
	input := parse_args()
}
