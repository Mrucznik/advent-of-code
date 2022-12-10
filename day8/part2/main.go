package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// time needed: 9min

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")
	heights := make([][]int, len(rows))
	viewScores := make([][]int, len(rows))
	for i, row := range rows {
		heights[i] = make([]int, len(row))
		viewScores[i] = make([]int, len(row))
		for j, height := range row {
			heights[i][j] = int(height)
		}
	}

	viewScore := 0
	mw, mh := len(rows), len(rows[0])
	for x := 0; x < mw; x++ {
		for y := 0; y < mh; y++ {
			vs := getViewScore(heights, x, y, mw, mh)
			viewScores[x][y] = vs
			if vs > viewScore {
				viewScore = vs
			}
		}
	}

	fmt.Println(viewScore)

}

func getViewScore(heights [][]int, myX, myY int, maxX, maxY int) int {
	myH := heights[myX][myY]

	var us, ds, ls, rs int

	// check up
	for y := myY - 1; y >= 0; y-- {
		us++
		if heights[myX][y] >= myH {
			break
		}
	}

	// check down
	for y := myY + 1; y < maxY; y++ {
		ds++
		if heights[myX][y] >= myH {
			break
		}
	}

	// check left
	for x := myX - 1; x >= 0; x-- {
		ls++
		if heights[x][myY] >= myH {
			break
		}
	}

	// check right
	for x := myX + 1; x < maxX; x++ {
		rs++
		if heights[x][myY] >= myH {
			break
		}
	}

	return us * ds * ls * rs
}
