package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// part 1: time needed 25min
// part 2: time needed 1h (i stuck on xd+yd >= 3)

//go:embed input.txt
var input string

func main() {
	simulation := NewSimulation()

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		raw := strings.Split(row, " ")
		move := raw[0][0]
		count, err := strconv.Atoi(raw[1])
		if err != nil {
			panic(err)
		}
		simulation.simulate(rune(move), count)
	}

	fmt.Println(simulation.getCount())
	fmt.Printf("%+v\n", simulation.visited)
}

type Simulation struct {
	visited map[string]bool
	tx      [9]int
	ty      [9]int
	hx      int
	hy      int
}

func NewSimulation() *Simulation {
	return &Simulation{visited: map[string]bool{"0x0": true}}
}

func (s *Simulation) simulate(move rune, count int) {
	for i := 0; i < count; i++ {
		s.move(move)
	}
}

func (s *Simulation) move(move rune) {
	// move head
	switch move {
	case 'L':
		s.hx--
	case 'R':
		s.hx++
	case 'U':
		s.hy--
	case 'D':
		s.hy++
	}

	for i := 0; i < 9; i++ {
		s.tailMove(i)
	}
}

func (s *Simulation) tailMove(nr int) {
	var hx, hy int
	if nr == 0 {
		hx = s.hx
		hy = s.hy
	} else {
		hx = s.tx[nr-1]
		hy = s.ty[nr-1]
	}
	tx := s.tx[nr]
	ty := s.ty[nr]

	xd := int(math.Abs(float64(tx - hx)))
	yd := int(math.Abs(float64(ty - hy)))

	if xd+yd > 4 {
		panic(fmt.Sprintf("%+v\nnr:%d xd:%d yd:%d\ntx: %d ty: %d hx: %d hy: %d", s.visited, nr, xd, yd, tx, ty, hx, hy))
	}

	if xd+yd >= 3 {
		// move towards head diagonally
		if hx > tx {
			tx++
		} else {
			tx--
		}
		if hy > ty {
			ty++
		} else {
			ty--
		}
	} else if xd+yd > 1 {
		if xd == 2 {
			// move towards head right/left
			if hx > tx {
				tx++
			} else {
				tx--
			}
		} else if yd == 2 {
			// move towards hea up/down
			if hy > ty {
				ty++
			} else {
				ty--
			}
		}
	}
	xd = int(math.Abs(float64(tx - hx)))
	yd = int(math.Abs(float64(ty - hy)))
	if xd > 1 || yd > 1 {
		panic(fmt.Sprintf("lol %+v\nnr:%d xd:%d yd:%d\ntx: %d ty: %d hx: %d hy: %d", s.visited, nr, xd, yd, tx, ty, hx, hy))
	}

	s.tx[nr] = tx
	s.ty[nr] = ty

	if nr == 8 {
		s.visited[fmt.Sprintf("%dx%d", tx, ty)] = true
	}
}

func (s *Simulation) getCount() int {
	c := 0
	for _, b := range s.visited {
		if b {
			c++
		}
	}
	return c
}

func getDist(tx, ty, hx, hy int) int {
	return int(math.Abs(float64(tx-hx)) + math.Abs(float64(ty-hy)))
}
