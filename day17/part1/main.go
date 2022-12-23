package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

const maxY, maxX = 100000, 7

var space = [maxY][maxX]uint8{}

func main() {

	rocks := []*Rock{
		NewRock([4][4]uint8{ // horizontal line
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}),
		NewRock([4][4]uint8{ // cross
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}),
		NewRock([4][4]uint8{ // L
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
		}),
		NewRock([4][4]uint8{ // vertical line
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		}),
		NewRock([4][4]uint8{ // cube
			{1, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}),
	}

	rows := strings.Split(input, "\n")
	for _, row := range rows {

	}
}
