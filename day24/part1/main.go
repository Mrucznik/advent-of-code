package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// blizards: reading 3:30
// part 1: 1h 11min +
// part 2:

type Elf struct {
	x, y   int
	mx, my int
	moves  int
}

func (e *Elf) move() {
	newx := (e.x + e.mx) % maxx
	if newx < 0 {
		newx += maxx
	}

	newy := (e.y + e.my) % maxy
	if newy < 0 {
		newx += maxy
	}

	e.x, e.y = newx, newy
	e.moves++
}

func (e *Elf) clone() []*Elf {
	return []*Elf{
		{x: e.x, y: e.y, mx: 0, my: 0, moves: e.moves},
		{x: e.x, y: e.y, mx: 0, my: 1, moves: e.moves},
		{x: e.x, y: e.y, mx: 1, my: 0, moves: e.moves},
		{x: e.x, y: e.y, mx: -1, my: 0, moves: e.moves},
		{x: e.x, y: e.y, mx: 0, my: -1, moves: e.moves},
	}
}

type Blizzard struct {
	x, y   int
	mx, my int
}

var maxx, maxy int

func NewBlizzard(x int, y int, mx int, my int) *Blizzard {
	return &Blizzard{x: x, y: y, mx: mx, my: my}
}

func (b *Blizzard) move() {
	b.x += b.mx
	b.y += b.my
}

func (b *Blizzard) collide(x, y int) bool {
	return b.x == x && b.y == y
}

func main() {
	blizzards := []*Blizzard{}
	rows := strings.Split(input, "\n")
	maxx = len(rows[0]) - 2
	maxy = len(rows) - 2
	for y, row := range rows[1 : len(rows)-1] {
		for x, char := range row[1 : len(row)-1] {
			var blizzard *Blizzard
			switch char {
			case '#', '.':
				continue
			case 'v':
				blizzard = NewBlizzard(x, y, 0, 1)
			case '>':
				blizzard = NewBlizzard(x, y, 1, 0)
			case '<':
				blizzard = NewBlizzard(x, y, -1, 0)
			case '^':
				blizzard = NewBlizzard(x, y, 0, -1)
			}

			blizzards = append(blizzards, blizzard)
		}
	}

	elves := map[int]*Elf{}
	won := make(chan *Elf)

	go func() {
		min := 100000000
		for elf := range won {
			if elf.moves < min {
				min = elf.moves
				fmt.Println("winner", elf.moves)
			}
		}
	}()

	totalElves := 0
	fmt.Println("stage 1")
	for i := 0; i < maxx*maxx*maxy*maxy; i++ {
		elves[totalElves] = &Elf{x: 0, y: 0, mx: 0, my: 0, moves: i}
		totalElves++

		for key, elf := range elves {
			for _, blizzard := range blizzards {
				blizzard.move()
			}

			if process(elf, blizzards) {
				delete(elves, key)
				continue
			}
			if elf.x == maxx-1 && elf.y == maxy-1 {
				// win
				delete(elves, key)
				won <- elf
				continue
			}

			for _, clone := range elf.clone() {
				if !process(clone, blizzards) {
					elves[totalElves] = clone
					totalElves++
				}
			}
		}
	}
}

func process(elf *Elf, blizzards []*Blizzard) bool {
	elf.move()

	if elf.x < 0 || elf.x >= maxx || elf.y < 0 || elf.y >= maxy {
		// kill
		return true
	}

	for _, blizzard := range blizzards {
		if blizzard.collide(elf.x, elf.y) {
			// kill
			return true
		}
	}
	return false
}
