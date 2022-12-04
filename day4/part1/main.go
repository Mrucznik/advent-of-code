package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// time needed: 13min

//go:embed input.txt
var input string

func main() {
	pairs := strings.Split(input, "\n")
	pairCount := getContainedPairCounts(pairs)
	fmt.Printf("result: %d\n", pairCount)
}

func getContainedPairCounts(pairs []string) int {
	pairCount := 0
	for _, pair := range pairs {
		ranges := strings.Split(pair, ",")

		min1, max1 := getElfRangeNumbers(ranges[0])
		min2, max2 := getElfRangeNumbers(ranges[1])

		if rangeContains(min1, max1, min2, max2) || rangeContains(min2, max2, min1, max1) {
			pairCount++
		}
	}
	return pairCount
}

func getElfRangeNumbers(elfRange string) (int, int) {
	numbers := strings.Split(elfRange, "-")
	min, _ := strconv.Atoi(numbers[0])

	max, _ := strconv.Atoi(numbers[1])
	return min, max
}

func rangeContains(min1, max1, min2, max2 int) bool {
	if min2 >= min1 && max2 <= max1 {
		return true
	}
	return false
}
