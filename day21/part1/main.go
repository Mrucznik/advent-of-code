package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// part 1 time needed: 12min 30s
// part 2 time needed:

//go:embed input.txt
var input string

func main() {
	chans := map[string]chan int{}

	for _, row := range strings.Split(input, "\n") {
		raw := strings.Split(row, ": ")
		monkey := raw[0]
		equation := raw[1]

		if val, err := strconv.Atoi(equation); err == nil {
			if _, ok := chans[monkey]; !ok {
				chans[monkey] = make(chan int)
			}

			go func(monke string) {
				chans[monkey] <- val
			}(monkey)
			continue
		} else {
			rawEq := strings.Split(equation, " ")
			m1, eq, m2 := rawEq[0], rawEq[1], rawEq[2]
			if _, ok := chans[monkey]; !ok {
				chans[monkey] = make(chan int)
			}
			if _, ok := chans[m1]; !ok {
				chans[m1] = make(chan int)
			}
			if _, ok := chans[m2]; !ok {
				chans[m2] = make(chan int)
			}

			go func(monke, m1, eq, m2 string) {
				a, b := <-chans[m1], <-chans[m2]
				switch eq {
				case "*":
					chans[monke] <- a * b
				case "/":
					chans[monke] <- a / b
				case "+":
					chans[monke] <- a + b
				case "-":
					chans[monke] <- a - b
				}
			}(monkey, m1, eq, m2)
		}
	}

	fmt.Println(<-chans["root"])
}
