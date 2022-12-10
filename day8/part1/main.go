package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// time needed: 25min

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")
	heights := make([][]int, len(rows))
	visible := make([][]bool, len(rows))
	for i, row := range rows {
		heights[i] = make([]int, len(row))
		visible[i] = make([]bool, len(row))
		for j, height := range row {
			heights[i][j] = int(height)
		}
	}

	var visibleCount int
	mw, mh := len(rows), len(rows[0])
	for x := 0; x < mw; x++ {
		for y := 0; y < mh; y++ {
			v := checkVisible(heights, x, y, mw, mh)
			visible[x][y] = v
			if v {
				visibleCount++
			}
		}
	}

	fmt.Println(visibleCount)

}

func checkVisible(heights [][]int, myX, myY int, maxX, maxY int) bool {
	myH := heights[myX][myY]

	if myX == 0 || myY == 0 || myX == maxX-1 || myY == maxY-1 {
		return true
	}

	checksFailed := 0
	// check up
	for y := myY - 1; y >= 0; y-- {
		if heights[myX][y] >= myH {
			checksFailed++
			fmt.Printf("x: %d y: %d visible up\n", myX, y)
			break
		}
	}

	// check down
	for y := myY + 1; y < maxY; y++ {
		if heights[myX][y] >= myH {
			checksFailed++
			break
		}
	}

	// check left
	for x := myX - 1; x >= 0; x-- {
		if heights[x][myY] >= myH {
			checksFailed++
			break
		}
	}

	// check right
	for x := myX + 1; x < maxX; x++ {
		if heights[x][myY] >= myH {
			checksFailed++
			break
		}
	}

	return checksFailed != 4
}
