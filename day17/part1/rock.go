package main

type Rock struct {
	x, y      int
	height int
	rockShape [4][4]uint8
}

func NewRock(rockShape [4][4]uint8) *Rock {
	h := 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if rockShape[y][x] == 1 {
				h++
				break
			}
		}
	}
	return &Rock{rockShape: rockShape, height: h}
}

func (r *Rock) spawn() {
	bottom := findSpaceBottom()
	r.y = bottom + 3 + r.height
	r.x = 2
}

var lastBottom int

func findSpaceBottom() int {
	bottom := lastBottom
	for y := bottom; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if space[y][x] == 1 {
				goto next
			}
		}
		return bottom
	next:
		bottom++
	}
	return bottom
}

func (r *Rock) move(left bool) {

}

func (r *Rock) fall() (stop bool) {
	if r.y == findSpaceBottom() {
		r.place()
	}
}

func (r *Rock) place() {
	for x := 0; x < ; x++ {
		
	}
}
