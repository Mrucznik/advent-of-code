package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

// day 13 part 1 time needed: 1h 25min

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")

	lines := make([][]interface{}, 0, len(rows))
	for _, line := range rows {
		if len(line) == 0 {
			continue
		}

		convertedLine, _ := convert(line)
		lines = append(lines, convertedLine[0].([]interface{}))
	}

	check(lines)

	rightOrderIndicesSum := 0

	for i := 0; i < len(lines); i += 2 {
		left := lines[i]
		right := lines[i+1]

		if comparison(left, right) == 1 {
			rightOrderIndicesSum += (i / 2) + 1
		}
	}

	fmt.Println(rightOrderIndicesSum)
}

func comparison(left, right interface{}) int {
	fmt.Printf("compare %v with %v\n", left, right)
	// convert
	var leftN, rightN int
	var leftL, rightL []interface{}
	switch val := left.(type) {
	case int:
		leftN = val
	case []interface{}:
		leftL = val
		if val == nil {
			val = []interface{}{}
		}
	default:
		panic("not a inter/number")
	}
	switch val := right.(type) {
	case int:
		rightN = val
	case []interface{}:
		rightL = val
		if val == nil {
			val = []interface{}{}
		}
	default:
		panic("not a inter/number")
	}

	if rightL == nil && leftL == nil {
		// both numbers
		if leftN < rightN {
			return 1
		} else if leftN == rightN {
			return 0
		} else {
			return -1
		}
	} else if rightL != nil && leftL != nil {
		// both lists
		max := int(math.Max(float64(len(leftL)), float64(len(rightL))))
		for i := 0; i < max; i++ {
			if i == len(leftL) {
				return 1
			}
			if i == len(rightL) {
				return -1
			}

			comp := comparison(leftL[i], rightL[i])
			if comp != 0 {
				return comp
			}
		}
		return 0
	} else {
		//fmt.Printf("L: %v R: %v\n", left, right)
		if rightL == nil {
			return comparison(left, []interface{}{rightN})
		} else if leftL == nil {
			return comparison([]interface{}{leftN}, right)
		}
		panic("nope")
	}
}
