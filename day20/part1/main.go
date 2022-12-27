package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// part 1 time needed: 3h because lack of L74 & L85
// part 2 time needed:

//go:embed input.txt
var input string

func main() {
	start := time.Now()

	rows := strings.Split(input, "\n")
	numbers := make([]int, len(rows))
	for i, row := range rows {
		var err error
		numbers[i], err = strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
	}

	l := len(numbers)

	//encrypt
	encrypted := encrypt(numbers)

	// coords
	zeroCor := 0
	for i, number := range encrypted {
		if number == 0 {
			zeroCor = i
			break
		}
	}
	fmt.Println(encrypted)

	fmt.Println("zero cor", zeroCor, encrypted[zeroCor])

	th1000 := encrypted[(zeroCor+1000)%l]
	th2000 := encrypted[(zeroCor+2000)%l]
	th3000 := encrypted[(zeroCor+3000)%l]
	fmt.Println(th1000, th2000, th3000)
	fmt.Println("sum", th1000+th2000+th3000)

	fmt.Println(time.Since(start))
}

func encrypt(numbers []int) []int {
	elements := list.New()

	var queue []*list.Element
	for _, number := range numbers {
		queue = append(queue, elements.PushBack(number))
	}

	for _, element := range queue {
		move := element.Value.(int)
		if move == 0 {
			continue
		}

		curr := element
		if move > 0 {
			for i := 0; i < move; i++ {
				curr = curr.Next()
				if curr == element {
					curr = curr.Next()
				}
				if curr == nil {
					curr = elements.Front()
				}
			}
			elements.MoveAfter(element, curr)
		} else {
			for i := 0; i < -move; i++ {
				curr = curr.Prev()
				if curr == element {
					curr = curr.Prev()
				}
				if curr == nil {
					curr = elements.Back()
				}
			}
			elements.MoveBefore(element, curr)
		}
	}

	arr := toArr(elements)
	return arr
}

func toArr(l *list.List) []int {
	arr := []int{}
	for e := l.Front(); e != nil; e = e.Next() {
		arr = append(arr, e.Value.(int))
	}
	return arr
}

// 1644 to low
