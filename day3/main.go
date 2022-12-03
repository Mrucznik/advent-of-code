package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rows := strings.Split(input, "\n")
	result := sumPrioritiesOfTheSameItemsInRucksackCompartments(rows)
	fmt.Println(result)
}

func sumPrioritiesOfTheSameItemsInRucksackCompartments(rucksack []string) int {
	itemSum := 0
	for _, ruck := range rucksack {
		item := extractTheSameItem(ruck)
		itemSum += itemSumToPriority(item)
	}

	return itemSum
}

func extractTheSameItem(ruck string) rune {
	items := map[rune]bool{}
	halfRuck := len(ruck) / 2
	i := 0
	for ; i < halfRuck; i++ {
		items[rune(ruck[i])] = true
	}
	for ; i < len(ruck); i++ {
		if items[rune(ruck[i])] {
			return rune(ruck[i])
		}
	}
	return -1000000000
}

func itemSumToPriority(item rune) int {
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	return -100000000
}
