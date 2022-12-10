package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// part1 time needed: 15min
// part 2 time needed: 18min

//go:embed input.txt
var input string

var X int

func main() {
	cycle := 0
	X = 1

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		raw := strings.Split(row, " ")
		cmd := raw[0]
		if cmd == "addx" {
			signalStrength(cycle)
			cycle++
			arg, err := strconv.Atoi(raw[1])
			if err != nil {
				panic(err)
			}
			signalStrength(cycle)
			cycle++

			X += arg
		} else {
			// noop
			signalStrength(cycle)
			cycle++
		}
	}
	fmt.Println()
	fmt.Println("result: ", result)
}

var result int

func signalStrength(c int) {
	if c%40 == 0 {
		fmt.Println("")
	}
	// draw
	draw := c % 40
	if X-1 == draw || X == draw || X+1 == draw {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

}
