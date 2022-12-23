package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const maxY, maxX = 100000, 7

var space = [maxX][maxY]uint8{}

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

	for i, rock := range rocks {
		fmt.Printf("rock %d width: %d height %d\n", i, rock.width, rock.height)
	}

	rows := strings.Split(input, "\n")
	allMoves := len(rows[0])
	moves := make([]bool, allMoves)
	for i, move := range rows[0] {
		moves[i] = rune(move) == '<'
	}

	currMove := 0
	currRock := 0
	fallen := 0
	rock := rocks[currRock]
	rock.spawn()
	for {
		rock.move(moves[currMove])
		currMove = (currMove + 1) % allMoves

		if rock.fall() {
			fallen++
			fmt.Printf("%d rock fallen %s\n", fallen, rock)
			currRock = (currRock + 1) % len(rocks)
			rock = rocks[currRock]
			rock.spawn()
		}

		if fallen == 2022 {
			break
		}
	}

	fmt.Println(findSpaceBottom())
}
