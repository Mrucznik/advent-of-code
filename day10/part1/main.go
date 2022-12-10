package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// part1 time needed: 15min

//go:embed input.txt
var input string

var X int

func main() {
	cycle := 1
	X = 1

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		raw := strings.Split(row, " ")
		cmd := raw[0]
		if cmd == "addx" {
			cycle++
			signalStrength(cycle)
			arg, err := strconv.Atoi(raw[1])
			if err != nil {
				panic(err)
			}
			cycle++

			X += arg
			signalStrength(cycle)

		} else {
			// noop
			cycle++
			signalStrength(cycle)
		}
	}
	fmt.Println("result: ", result)
}

var result int

func signalStrength(c int) {
	if c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220 {
		result += X * c
		fmt.Printf("X: %d c: %d r: %d\n", X, c, X*c)
	}
}
