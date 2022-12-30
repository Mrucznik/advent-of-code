package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// part 1 time needed: 26min + ~30min
// part 2 time needed: ~5min

//go:embed input.txt
var input string

var space = [1000][1000]*Elf{}

type Elf struct {
	checksQueue []int
}

func NewElf() *Elf {
	return &Elf{
		checksQueue: []int{0, 1, 2, 3},
	}
}

type Pos struct {
	x, y int
}

func main() {
	rows := strings.Split(input, "\n")
	add := 1000/2 - len(rows)/2
	for y, row := range rows {
		for x, char := range row {
			if char == '#' {
				space[add+y][add+x] = NewElf()
			}
		}
	}

	for i := 0; i < 10000; i++ {
		if move() {
			fmt.Println("odp", i+1)
			break
		}
	}

	maxx, maxy, minx, miny := findSmallestRect()
	fmt.Println(countEmpty(maxx, maxy, minx, miny))
}

func countEmpty(maxx, maxy, minx, miny int) (count int) {
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if space[y][x] == nil {
				count++
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	return
}

func findSmallestRect() (int, int, int, int) {
	var maxx, maxy int
	minx, miny := 1000, 1000

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if space[y][x] != nil {
				if x > maxx {
					maxx = x
				}
				if x < minx {
					minx = x
				}
				if y > maxy {
					maxy = y
				}
				if y < miny {
					miny = y
				}
			}
		}
	}

	return maxx, maxy, minx, miny
}

func move() bool {
	oldPos := map[*Elf]Pos{}
	newSpace := [1000][1000][]*Elf{}
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if space[y][x] != nil {
				newx, newy := checkPosToMove(x, y)
				oldPos[space[y][x]] = Pos{x, y}
				newSpace[newy][newx] = append(newSpace[newy][newx], space[y][x])
			}
		}
	}

	newSpace2 := [1000][1000]*Elf{}
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if len(newSpace[y][x]) > 1 {
				for _, elf := range newSpace[y][x] {
					old := oldPos[elf]
					newSpace2[old.y][old.x] = elf
					elf.checksQueue = append(elf.checksQueue[1:], elf.checksQueue[0])
				}
			} else if len(newSpace[y][x]) == 1 {
				elf := newSpace[y][x][0]
				newSpace2[y][x] = elf
				elf.checksQueue = append(elf.checksQueue[1:], elf.checksQueue[0])
			}
		}
	}

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if space[y][x] != newSpace2[y][x] {
				space = newSpace2
				return false
			}
		}
	}
	space = newSpace2
	return true
}

func checkPosToMove(cx, cy int) (int, int) {
	var occN, occS, occW, occE bool

	// check elves around
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}

			if space[cy+y][cx+x] != nil {
				if y == 1 {
					occS = true
				} else if y == -1 {
					occN = true
				}

				if x == 1 {
					occE = true
				} else if x == -1 {
					occW = true
				}
			}
		}
	}

	if occN || occS || occW || occE {
		checks := []bool{occN, occS, occW, occE}
		px, py := proposeMove(cx, cy, checks)
		return px, py
	} else {
		return cx, cy
	}
}

func proposeMove(x int, y int, checks []bool) (int, int) {
	elf := space[y][x]

	for _, idx := range elf.checksQueue {
		if !checks[idx] {
			return checkToMove(x, y, idx)
		}
	}
	return x, y
}

func checkToMove(x, y, check int) (int, int) {
	switch check {
	case 0:
		return x, y - 1
	case 1:
		return x, y + 1
	case 2:
		return x - 1, y
	case 3:
		return x + 1, y
	}
	panic("no")
}
