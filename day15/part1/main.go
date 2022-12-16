package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// part 1 took 34min

// part 2 took 1h 24min

//go:embed input.txt
var input string

type Point struct {
	x, y int
}
type Sensor struct {
	Point
	beaconDistance int
	distX, distY   int
}

func main() {
	re := regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")

	beacons := []*Point{}
	sensors := []*Sensor{}

	max := 4000000

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		raw := re.FindStringSubmatch(row)

		sx, _ := strconv.Atoi(raw[1])
		sy, _ := strconv.Atoi(raw[2])
		sp := Point{sx, sy}

		bx, _ := strconv.Atoi(raw[3])
		by, _ := strconv.Atoi(raw[4])
		bp := Point{bx, by}

		distX := Abs(bp.x - sp.x)
		distY := Abs(bp.y - sp.y)
		dist := distY + distX
		sensors = append(sensors, &Sensor{sp, dist, distX, distY})
		beacons = append(beacons, &bp)
	}

	pointsToCheck := make(chan Point, 100000)

	wg := &sync.WaitGroup{}
	for _, sensor := range sensors {
		wg.Add(2)
		go func(sensor *Sensor) {
			for x := (-sensor.beaconDistance) - 1; x < sensor.beaconDistance+1; x++ {
				y := Abs(sensor.beaconDistance + 1 - x)

				point := Point{sensor.x + x, sensor.y - y}
				if point.x < 0 || point.x > max || point.y < 0 || point.y > max {
					continue
				}
				pointsToCheck <- point
			}
			wg.Done()
		}(sensor)

		go func(sensor *Sensor) {
			for x := (-sensor.beaconDistance) - 1; x < sensor.beaconDistance+1; x++ {
				y := Abs(sensor.beaconDistance + 1 - x)

				point := Point{sensor.x + x, sensor.y + y}
				if point.x < 0 || point.x > max || point.y < 0 || point.y > max {
					continue
				}
				pointsToCheck <- point
			}
			wg.Done()
		}(sensor)
	}

	fmt.Println("checking")
	for point := range pointsToCheck {
		otherSensorOverlap := false
		for _, sensor := range sensors {
			dist := Abs(sensor.y-point.y) + Abs(sensor.x-point.x)

			if dist <= sensor.beaconDistance {
				otherSensorOverlap = true
				break
			}
		}

		if !otherSensorOverlap {
			beaconOverlap := false
			for _, beacon := range beacons {
				if point.x == beacon.x && point.y == beacon.y {
					beaconOverlap = true
					break
				}
			}

			if !beaconOverlap {
				fmt.Println("wynik: ", point.x*4000000+point.y)
				return
			}
		}
	}
	fmt.Println("1")
	wg.Wait()
	fmt.Println("nothing")
}
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
