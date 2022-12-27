package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// part 1 time needed: 12min 30s
// part 2 time needed:

//go:embed input.txt
var input string

func main() {
	chans := &sync.Map{}

	for _, row := range strings.Split(input, "\n") {
		raw := strings.Split(row, ": ")
		monkey := raw[0]
		equation := raw[1]

		if val, err := strconv.Atoi(equation); err == nil {
			if _, ok := chans.Load(monkey); !ok {
				chans.Store(monkey, make(chan int))
			}

			go func(monke string) {
				for {
					c, _ := chans.Load(monke)
					c.(chan int) <- val
				}
			}(monkey)
			continue
		} else {
			rawEq := strings.Split(equation, " ")
			m1, eq, m2 := rawEq[0], rawEq[1], rawEq[2]
			if _, ok := chans.Load(monkey); !ok {
				chans.Store(monkey, make(chan int))
			}
			if _, ok := chans.Load(m1); !ok {
				chans.Store(m1, make(chan int))
			}
			if _, ok := chans.Load(m2); !ok {
				chans.Store(m2, make(chan int))
			}

			go func(monke, m1, eq, m2 string) {
				for {
					c1, _ := chans.Load(m1)
					c2, _ := chans.Load(m2)

					a, b := <-c1.(chan int), <-c2.(chan int)
					switch eq {
					case "*":
						c, _ := chans.Load(monke)
						c.(chan int) <- a * b
					case "/":
						c, _ := chans.Load(monke)
						c.(chan int) <- a / b
					case "+":
						c, _ := chans.Load(monke)
						c.(chan int) <- a + b
					case "-":
						c, _ := chans.Load(monke)
						c.(chan int) <- a - b
					case "=":
						if a == b {
							c, _ := chans.Load(monke)
							c.(chan int) <- 1
						} else {
							c, _ := chans.Load(monke)
							c.(chan int) <- 0
						}
					}
				}
			}(monkey, m1, eq, m2)
		}
	}

	go func() {
		i := 100000
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("aaa", i)
			}
		}()
		for {
			c, _ := chans.Load("humn")
			c.(chan int) <- i
			i++
		}

	}()

	i := 0
	for {
		c, _ := chans.Load("humn")
		rc, _ := chans.Load("root")
		if <-rc.(chan int) == 1 {
			fmt.Println("wynik", i)
			close(c.(chan int))
			break
		}
		i++
	}
}
