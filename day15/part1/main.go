package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// part 1 took 34min

// part 2 took

//go:embed input.txt
var input string

type Point struct {
	x, y int
}
type Sensor struct {
	Point
	beaconDistance int
}

func main() {
	re := regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")

	beacons := []Point{}
	sensors := []Sensor{}

	maxX := 0
	minX := 0

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		raw := re.FindStringSubmatch(row)

		sx, _ := strconv.Atoi(raw[1])
		sy, _ := strconv.Atoi(raw[2])
		sp := Point{sx, sy}

		bx, _ := strconv.Atoi(raw[3])
		by, _ := strconv.Atoi(raw[4])
		bp := Point{bx, by}

		dist := distance(sp, bp)
		sensors = append(sensors, Sensor{sp, dist})
		beacons = append(beacons, bp)

		if sx+dist > maxX {
			maxX = sx + dist + 1
		}

		if sx-dist < minX {
			minX = sx - dist - 1
		}
	}

	y := 2000000
	pos := 0
	for x := minX - 10; x < maxX+10; x++ {
		for _, sensor := range sensors {
			checkPoint := Point{x, y}
			dist := distance(checkPoint, sensor.Point)

			for _, beacon := range beacons {
				if checkPoint.x == beacon.x && checkPoint.y == beacon.y {
					goto next
				}
			}

			if dist <= sensor.beaconDistance {
				pos++
				break
			}
		next:
		}
	}

	fmt.Println(pos)
}

func distance(a Point, b Point) int {
	return int(math.Abs(float64(b.x-a.x))) + int(math.Abs(float64(b.y-a.y)))
}
