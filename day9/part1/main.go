package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// time needed 25min

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
}

type Simulation struct {
	visited map[string]bool
	tx      int
	ty      int
	hx      int
	hy      int
}

func NewSimulation() *Simulation {
	return &Simulation{visited: map[string]bool{"0-0": true}}
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

	xd := int(math.Abs(float64(s.tx - s.hx)))
	yd := int(math.Abs(float64(s.ty - s.hy)))

	if xd+yd == 3 {
		// move towards head diagonally
		if s.hx > s.tx {
			s.tx++
		} else {
			s.tx--
		}
		if s.hy > s.ty {
			s.ty++
		} else {
			s.ty--
		}
	} else if xd+yd > 1 {
		if xd == 2 {
			// move towards head right/left
			if s.hx > s.tx {
				s.tx++
			} else {
				s.tx--
			}
		} else if yd == 2 {
			// move towards hea up/down
			if s.hy > s.ty {
				s.ty++
			} else {
				s.ty--
			}
		}
	}

	// mark
	s.visited[fmt.Sprintf("%d-%d", s.tx, s.ty)] = true
}

func (s Simulation) getCount() int {
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
