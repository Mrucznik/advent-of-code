package main

import "fmt"

const rockGridSize = 4

type RockGrid [rockGridSize][rockGridSize]uint8

type Rock struct {
	x, y      int
	height    int
	width     int
	rockShape RockGrid
}

func NewRock(rockShape RockGrid) *Rock {
	h := 0
	for y := 0; y < rockGridSize; y++ {
		for x := 0; x < rockGridSize; x++ {
			if rockShape[y][x] == 1 {
				h++
				break
			}
		}
	}

	w := 0
	for x := 0; x < rockGridSize; x++ {
		for y := 0; y < rockGridSize; y++ {
			if rockShape[y][x] == 1 {
				w++
				break
			}
		}
	}
	return &Rock{rockShape: rockShape, height: h, width: w}
}

func (r *Rock) spawn() {
	bottom := findSpaceBottom()
	r.y = bottom + 3 + r.height - 1
	r.x = 2
}

var lastBottom int

func findSpaceBottom() int {
	for y := lastBottom; y < maxY; y++ {
		nextRow := false
		for x := 0; x < maxX; x++ {
			if space[x][y] == 1 {
				nextRow = true
				break
			}
		}

		if !nextRow {
			lastBottom = y
			return y
		}
	}
	panic("no")
}

func (r *Rock) move(left bool) {
	// collide with walls
	if r.x <= 0 && left {
		return
	}

	if r.x+r.width >= 7 && !left {
		return
	}

	if left {
		r.x--
		if r.collide() {
			r.x++
		}
	} else {
		r.x++
		if r.collide() {
			r.x--
		}
	}
}

func (r *Rock) fall() (stop bool) {
	if r.y-r.height+1 <= 0 {
		r.place()
		return true
	}
	r.y--
	if r.collide() {
		r.y++
		r.place()
		return true
	}
	return false
}

func (r *Rock) place() {
	for x := 0; x < rockGridSize; x++ {
		for y := 0; y < rockGridSize; y++ {
			if r.rockShape[y][x] == 1 {
				space[r.x+x][r.y-y] = 1
			}
		}
	}
}

func (r *Rock) collide() bool {
	for x := 0; x < rockGridSize; x++ {
		for y := 0; y < rockGridSize; y++ {
			if r.rockShape[y][x] == 1 && space[r.x+x][r.y-y] == 1 {
				return true
			}
		}
	}
	return false
}

func (r *Rock) String() string {
	return fmt.Sprintf("rock x: %d y: %d width: %d height: %d", r.x, r.y, r.width, r.height)
}
