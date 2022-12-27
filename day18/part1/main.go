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

var space = [100][100][100]int{}

func main() {
	rows := strings.Split(input, "\n")

	droplets := map[Pos]struct{}{}
	for _, row := range rows {
		coords := strings.Split(row, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		space[x][y][z] = 1
		droplets[Pos{x, y, z}] = struct{}{}
	}

	// mark outsides
	markOutsides(Pos{})

	sides := 0
	for pos := range droplets {
		// count sides
		sides += sidesCount(pos)
	}

	fmt.Println(sides)
}

func markOutsides(pos Pos) {
	if pos.x < 0 || pos.y < 0 || pos.z < 0 || pos.x >= 100 || pos.y >= 100 || pos.z >= 100 {
		return
	}

	if space[pos.x][pos.y][pos.z] != 0 {
		return
	}

	space[pos.x][pos.y][pos.z] = -1
	markOutsides(Pos{x: pos.x, y: pos.y, z: pos.z + 1})
	markOutsides(Pos{x: pos.x, y: pos.y + 1, z: pos.z})
	markOutsides(Pos{x: pos.x + 1, y: pos.y, z: pos.z})
	markOutsides(Pos{x: pos.x - 1, y: pos.y, z: pos.z})
	markOutsides(Pos{x: pos.x, y: pos.y - 1, z: pos.z})
	markOutsides(Pos{x: pos.x, y: pos.y, z: pos.z - 1})
}

func sidesCount(pos Pos) (sides int) {
	sides += checkSide(pos, 0, 0, 1)
	sides += checkSide(pos, 0, 1, 0)
	sides += checkSide(pos, 1, 0, 0)
	sides += checkSide(pos, 0, 0, -1)
	sides += checkSide(pos, 0, -1, 0)
	sides += checkSide(pos, -1, 0, 0)
	return sides
}

func checkSide(pos Pos, tx, ty, tz int) int {
	if pos.x+tx < 0 || pos.y+ty < 0 || pos.z+tz < 0 {
		return 1
	}

	if space[pos.x+tx][pos.y+ty][pos.z+tz] != -1 {
		return 0
	}

	return 1
}
