package main

import (
	"bytes"
	_ "embed"
	"encoding/binary"
	"fmt"
	"github.com/cespare/xxhash"
	"strings"
)

//go:embed input.txt
var input string

const maxY = 2_000_000

var space = [maxY]byte{}

func main() {
	rocks := [5][4]byte{
		{
			0b0_0000000,
			0b0_0000000,
			0b0_0000000,
			0b0_1111000 >> 2,
		},
		{
			0b0_0000000,
			0b0_0100000 >> 2,
			0b0_1110000 >> 2,
			0b0_0100000 >> 2,
		},
		{
			0b0_0000000,
			0b0_0010000 >> 2,
			0b0_0010000 >> 2,
			0b0_1110000 >> 2,
		},
		{
			0b0_1000000 >> 2,
			0b0_1000000 >> 2,
			0b0_1000000 >> 2,
			0b0_1000000 >> 2,
		},
		{
			0b0_0000000,
			0b0_0000000,
			0b0_1100000 >> 2,
			0b0_1100000 >> 2,
		},
	}

	const toFall = maxY / 2
	fallen := 0

	rows := strings.Split(input, "\n")
	moves := rows[0]
	allMoves := len(moves)

	// spawn
	y := 6
	currMove := 0
	currRock := 0
	rock := rocks[currRock]

	snapshots := map[uint64]int{}
	var snapshotsSlice []uint64
	repeatedFrom := uint64(0)
	endLength := 0
	fillupLength := make([]int, 100000)
	ff := 0
	for fallen != toFall {
		// move
		move := moves[currMove%allMoves]
		currMove++
		if move == '<' {
			rock = moveLeft(rock, y)
		} else {
			rock = moveRight(rock, y)
		}

		// fall
		stop := fall(rock, y)
		if stop {
			// place
			place(rock, y)

			// spawn
			bottom := findSpaceBottom(y - 3)
			y = bottom + 6
			currRock++
			fallen++
			rock = rocks[currRock%len(rocks)]

			if endLength == 0 {
				ss := snapshot(currRock%len(rocks), currMove%allMoves, y)
				if _, ok := snapshots[ss]; ok {
					repeatedFrom = ss
					snapshotsSlice = append(snapshotsSlice, ss)
					endLength = bottom
				} else {
					snapshots[ss] = bottom
					snapshotsSlice = append(snapshotsSlice, ss)
				}
			} else {
				ff++
				fillupLength[ff] = bottom - endLength
			}
			if ff == 10000 {
				break
			}
		} else {
			y--
		}
	}

	for i, key := range snapshotsSlice {
		fmt.Println(i, key)
	}

	beforeCycleSteps := 0
	cycleLength := 0
	for i, key := range snapshotsSlice {
		if key == repeatedFrom {
			beforeCycleSteps = i
			cycleLength = len(snapshotsSlice) - i - 1
			break
		}
	}

	result := 0
	// 1. wysokość od rozpoczęcia cyklu do upadnięcia 1 z cyklu
	result += snapshots[snapshotsSlice[beforeCycleSteps]]
	fmt.Println("before cycle steps", beforeCycleSteps)
	fmt.Println("result", result)

	// 2. wysokość cyklu (od 1 z cyklu odjąć ostatni z cyklu)
	cycleHeight := endLength - snapshots[snapshotsSlice[beforeCycleSteps]]
	fmt.Println("cycle height", cycleHeight)
	fmt.Println("cycle length", cycleLength)

	// 3. ile razy cykl tyle powtórzyć
	repeatTimes := (1000000000000 - beforeCycleSteps) / cycleLength
	dopelnienie := (1000000000000 - beforeCycleSteps) % cycleLength
	fmt.Println("diff", 1000000000000-beforeCycleSteps)
	fmt.Println("repeat", repeatTimes)
	fmt.Println("dop", dopelnienie)
	result += repeatTimes * cycleHeight
	fmt.Println("result", result)

	// 4. wysokość po dopełnieniu do cyklu
	dopHeight := fillupLength[dopelnienie]
	fmt.Println("dop h", dopHeight)
	result += dopHeight
	fmt.Println("result", result)

	// not 1533141210385
	// not 1533141210386
	// not 1532564841509
	// not 1532564841510
	// not 1534870317014
	// not 1534870317013
	// not 1534870317012
	// not 1534870317011
	// not 1534870317010
}

func snapshot(currRock int, currMove int, y int) uint64 {
	sb := xxhash.New()
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.BigEndian, int64(currRock))
	binary.Write(buff, binary.BigEndian, int64(currMove))
	sb.Write(buff.Bytes())

	from := 10091 * 2
	if y < from {
		from = y
	}
	sb.Write(space[from:y])

	return sb.Sum64()
}

func place(rock [4]byte, y int) {
	for i := 0; i < len(rock); i++ {
		space[y-i] |= rock[i]
	}
}

func fall(rock [4]byte, y int) bool {
	for i := 0; i < len(rock); i++ {
		spaceY := y - 1 - i
		if spaceY < 0 {
			return true
		}
		if rock[i]&space[spaceY] != 0 {
			return true
		}
	}
	return false
}

func moveLeft(rock [4]byte, y int) [4]byte {
	nextRock := rock
	for i := 0; i < len(nextRock); i++ {
		nextRock[i] <<= 1
		if nextRock[i]&0b10000000 == 0b10000000 {
			return rock
		}

		if space[y-i]&nextRock[i] != 0 {
			return rock
		}
	}
	return nextRock
}

func moveRight(rock [4]byte, y int) [4]byte {
	nextRock := rock
	for i := 0; i < len(nextRock); i++ {
		if nextRock[i]&0b00000001 == 0b00000001 {
			return rock
		}
		nextRock[i] >>= 1

		if space[y-i]&nextRock[i] != 0 {
			return rock
		}
	}
	return nextRock
}

func findSpaceBottom(beginY int) int {
	for y := beginY; y < maxY; y++ {
		if space[y]&0b01111111 == 0 {
			return y
		}
	}
	return 0
}
