package main

import (
	_ "embed"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"math"
	"strings"
)

// time needed: 1h 44min

// time needed: 5min 30s

//go:embed input.txt
var input string

type xy struct {
	x int
	y int
}

var rows []string

func main() {
	rows = strings.Split(input, "\n")

	SID := []string{}
	EID := ""

	graph := dijkstra.NewGraph()
	// find begginning nodes
	for y, row := range rows {
		for x, letter := range row {
			id := getID(x, y)

			if letter == 'a' {
				SID = append(SID, id)
			} else if letter == 'E' {
				EID = id
			}

			graph.AddMappedVertex(id)
		}
	}

	// add arcs
	for myY, row := range rows {
		for myX, letter := range row {
			// right and down
			if myX+1 < len(row) {
				myID, neighbourID := getID(myX, myY), getID(myX+1, myY)
				if isAccessible(letter, rune(rows[myY][myX+1])) {
					checkErr(graph.AddMappedArc(myID, neighbourID, 1))
				}
				if isAccessible(rune(rows[myY][myX+1]), letter) {
					checkErr(graph.AddMappedArc(neighbourID, myID, 1))
				}
			}
			if myY+1 < len(rows) {
				myID, neighbourID := getID(myX, myY), getID(myX, myY+1)
				if isAccessible(letter, rune(rows[myY+1][myX])) {
					checkErr(graph.AddMappedArc(myID, neighbourID, 1))
				}
				if isAccessible(rune(rows[myY+1][myX]), letter) {
					checkErr(graph.AddMappedArc(neighbourID, myID, 1))
				}
			}
		}
	}

	var min int64 = math.MaxInt64
	for _, rawSID := range SID {
		sid, err := graph.GetMapping(rawSID)
		if err != nil {
			panic(err)
		}
		eid, err := graph.GetMapping(EID)
		if err != nil {
			panic(err)
		}

		shortest, err := graph.Shortest(sid, eid)
		if err != nil {
			continue
		}

		if shortest.Distance < min {
			min = shortest.Distance
		}
	}

	fmt.Printf("%+v\n", min)
	//for _, v := range shortest.Path {
	//	fmt.Println(graph.GetMapped(v))
	//}
}

func isAccessible(a, b rune) bool {
	if a == 'S' {
		a = 'a' - 1
	}
	if b == 'S' {
		b = 'a' - 1
	}

	if a == 'E' {
		a = 'z' + 1
	}
	if b == 'E' {
		b = 'z' + 1
	}

	x := b - a
	if x > 1 {
		return false
	}
	return true
}

func getID(x, y int) string {
	return fmt.Sprintf("%dx%d_%c", x, y, rows[y][x])
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
