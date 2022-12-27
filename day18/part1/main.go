package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Pos struct {
	x, y, z int
}

var space = [100][100][100]bool{}

func main() {
	rows := strings.Split(input, "\n")

	droplets := map[Pos]struct{}{}
	for _, row := range rows {
		coords := strings.Split(row, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		space[x][y][z] = true
		droplets[Pos{x, y, z}] = struct{}{}
	}

	sides := 0
	for pos := range droplets {
		// count sides
		sides += checkSide(pos, 0, 0, 1)
		sides += checkSide(pos, 0, 1, 0)
		sides += checkSide(pos, 1, 0, 0)
		sides += checkSide(pos, 0, 0, -1)
		sides += checkSide(pos, 0, -1, 0)
		sides += checkSide(pos, -1, 0, 0)
	}

	fmt.Println(sides)
}

func checkSide(pos Pos, tx, ty, tz int) int {
	if pos.x+tx < 0 || pos.y+ty < 0 || pos.z+tz < 0 {
		return 1
	}

	if space[pos.x+tx][pos.y+ty][pos.z+tz] {
		return 0
	}
	return 1
}
