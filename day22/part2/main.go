package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// part 1 took 2h (last letter in input was not read)
// part 2 took

//go:embed input.txt
var input string

type Ja struct {
	x, y, facing int
}

const maxX, maxY = 150, 200

func (ja *Ja) nextCoord(times int) (int, int) {
	if ja.facing == 0 { // right
		return (ja.x + times) % maxX, ja.y
	} else if ja.facing == 1 { // down
		return ja.x, (ja.y + times) % maxY
	} else if ja.facing == 2 { // left
		x := ja.x - times
		if x < 0 {
			x += maxX
		}
		return x, ja.y
	} else if ja.facing == 3 { // up
		y := ja.y - times
		if y < 0 {
			y += maxY
		}
		return ja.x, y
	}
	panic("no")
}

func (ja *Ja) L() {
	if ja.facing == 0 {
		ja.facing = 3
		return
	}
	ja.facing = ja.facing - 1
}

func (ja *Ja) R() {
	if ja.facing == 3 {
		ja.facing = 0
		return
	}
	ja.facing = ja.facing + 1
}

func (ja *Ja) move(times int) {
	spaces := 0
	for i := 0; i < times; {
		x, y := ja.nextCoord(1 + i + spaces)
		tile := space[y][x]
		if tile == ' ' {
			spaces++
			continue
		}

		if tile == '#' {
			// back
			for j := 0; j <= spaces; j++ {
				x, y = ja.nextCoord(i + spaces - j)
				if space[y][x] != ' ' {
					ja.x, ja.y = x, y
					return
				}
			}
		}

		i++
	}
	ja.x, ja.y = ja.nextCoord(times + spaces)
}

func (ja Ja) String() string {
	return fmt.Sprintf("x: %d, y: %d, facing: %d", ja.x, ja.y, ja.facing)
}

var space = [maxY][maxX]rune{}

func main() {
	rows := strings.Split(input, "\n")

	// load map
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if x >= len(rows[y]) {
				space[y][x] = ' '
			} else {
				space[y][x] = rune(rows[y][x])
			}
		}
	}

	// form cubes

	// sequences
	sequences := rows[maxY+1]
	fmt.Println(sequences)

	ja := Ja{0, 0, 0}

	re := regexp.MustCompile("(\\d+|[LR])")
	matches := re.FindAllStringSubmatch(sequences, -1)
	for _, match := range matches {
		steps, err := strconv.Atoi(match[1])
		if err != nil {
			turn := match[1]
			if turn == "L" {
				ja.L()
			} else {
				ja.R()
			}
		} else {
			// do the move
			ja.move(steps)
		}
	}

	fmt.Println(1000*(ja.y+1) + 4*(ja.x+1) + ja.facing)
}
