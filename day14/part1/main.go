package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// part 1 time needed: 53min 18s

// part 2 time needed ~5min

//go:embed input.txt
var input string

type Point struct {
	x int
	y int
}

type Tile int

const (
	AIR Tile = iota
	ROCK
	SAND
)

var world = [10000][10000]Tile{}
var maxY int

func main() {

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		rawPoints := strings.Split(row, " -> ")
		points := make([]Point, len(rawPoints))
		for i, pointData := range rawPoints {
			coords := strings.Split(pointData, ",")
			points[i].x, _ = strconv.Atoi(coords[0])
			points[i].y, _ = strconv.Atoi(coords[1])
		}

		var begin Point
		begin = points[0]
		for _, point := range points[1:] {
			drawRockLine(begin, point)
			begin = point
			if maxY < point.y {
				maxY = point.y
			}
		}
	}
	maxY += 2
	drawRockLine(Point{0, maxY}, Point{10000 - 1, maxY})

	sandPoint := Point{x: 500, y: 0}
	fmt.Println(pourSand(sandPoint))
}

func draw() {
	for y := 0; y < 10; y++ {
		for x := 494; x < 504; x++ {
			switch world[x][y] {
			case AIR:
				fmt.Print(".")
			case ROCK:
				fmt.Print("#")
			case SAND:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

func drawRockLine(p1 Point, p2 Point) {
	fmt.Printf("draw line from %v to %v\n", p1, p2)

	if p1.x != p2.x {
		var fromX, toX int
		if p1.x > p2.x {
			toX = p1.x
			fromX = p2.x
		} else {
			toX = p2.x
			fromX = p1.x
		}

		for i := fromX; i <= toX; i++ {
			world[i][p1.y] = ROCK
		}
	}

	if p1.y != p2.y {
		var fromY, toY int
		if p1.y > p2.y {
			toY = p1.y
			fromY = p2.y
		} else {
			toY = p2.y
			fromY = p1.y
		}

		for i := fromY; i <= toY; i++ {
			world[p1.x][i] = ROCK
		}
	}
}

func pourSand(sandPoint Point) int {
	sandUnits := 0
	for {
	begin:
		//spawn sand & fall
		sand := Point{sandPoint.x, sandPoint.y}
		sandUnits++
		for i := 1; i < 1000; i++ {
			var stop bool
			sand, stop = fallSand(sand)
			if stop {
				if sandPoint == sand {
					return sandUnits
				}
				world[sand.x][sand.y] = SAND
				//draw()
				goto begin
			}
			if i == 900 {
				return sandUnits - 1
			}
		}
	}
}

func fallSand(sandPoint Point) (Point, bool) {
	if world[sandPoint.x][sandPoint.y+1] == AIR {
		return Point{sandPoint.x, sandPoint.y + 1}, false
	} else if world[sandPoint.x-1][sandPoint.y+1] == AIR {
		return Point{sandPoint.x - 1, sandPoint.y + 1}, false
	} else if world[sandPoint.x+1][sandPoint.y+1] == AIR {
		return Point{sandPoint.x + 1, sandPoint.y + 1}, false
	} else {
		return Point{sandPoint.x, sandPoint.y}, true
	}
}
