package main

import (
	_ "embed"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")

	SID := 0
	EID := 0

	graph := dijkstra.NewGraph()
	// find begginning nodes
	for y, row := range rows {
		for x, letter := range row {
			id := getID(x, y)
			if letter == 'S' {
				SID = id
			}
			if letter == 'E' {
				EID = id
			}
			graph.AddVertex(id)
		}
	}

	// add arcs
	for myY, row := range rows {
		if myY+1 >= len(rows) {
			break
		}
		for myX, letter := range row {
			if myX+1 >= len(row) {
				break
			}

			// right and down
			if isAccessible(letter, rune(rows[myY][myX+1])) {
				fmt.Printf("x %c %c\n", letter, rune(rows[myY][myX+1]))
				myID, neighbourID := getID(myX, myY), getID(myX+1, myY)
				checkErr(graph.AddArc(myID, neighbourID, 1))
				checkErr(graph.AddArc(neighbourID, myID, 1))
			}
			if isAccessible(letter, rune(rows[myY+1][myX])) {
				fmt.Printf("y %c %c\n", letter, rune(rows[myY][myX+1]))
				myID, neighbourID := getID(myX, myY), getID(myX, myY+1)
				checkErr(graph.AddArc(myID, neighbourID, 1))
				checkErr(graph.AddArc(neighbourID, myID, 1))
			}
		}
	}

	shortest, err := graph.Shortest(SID, EID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", shortest)
}

func isAccessible(a, b rune) bool {
	if a == 'S' || a == 'E' || b == 'S' || b == 'E' {
		return true
	}

	x := a - b
	if x < -1 || x > 1 {
		return false
	}
	return true
}

func getID(x, y int) int {
	return x*1000 + y
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
