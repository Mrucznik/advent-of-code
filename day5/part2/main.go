package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// time needed: 2-3min :(

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")
	fmt.Println(doRearrangement(rows))
}

// test
//const N = 3
//const M = 3
// prod
const N = 8
const M = 9

func doRearrangement(rows []string) string {
	stacks := createStacks(rows)

	for _, row := range rows[N+2:] {
		quantity, from, to := ParseMove(row)

		// move
		//fmt.Printf("%d %d %d | %s | %s %d\n", quantity, from, to, row, string(stacks[from]), len(stacks[from]))

		len := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][len-quantity:]...)
		stacks[from] = stacks[from][:len-quantity]
	}

	result := strings.Builder{}
	for _, stack := range stacks {
		result.WriteRune(stack[len(stack)-1])
	}
	return result.String()
}

func createStacks(rows []string) [M][]rune {
	stacks := [M][]rune{}
	for i := N - 1; i >= 0; i-- {
		row := rows[i]
		for j := 0; j < M; j++ {
			letter := rune(row[1+j*4])
			if letter != ' ' {
				stacks[j] = append(stacks[j], letter)
			}
		}
	}
	return stacks
}

func ParseMove(move string) (int, int, int) {
	if len(move) == 0 {
		return 0, 0, 0
	}

	re := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	matches := re.FindStringSubmatch(move)
	if len(matches) != 4 {
		fmt.Println(move)
	}
	quantity, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])

	from = from - 1
	to = to - 1
	return quantity, from, to
}
