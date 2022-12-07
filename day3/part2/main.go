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
	for i := 0; i < len(rucksack); i += 3 {
		item := extractGroupItem(rucksack[i], rucksack[i+1], rucksack[i+2])
		itemSum += itemSumToPriority(item)
		fmt.Printf("Item: %c\n", item)
	}

	return itemSum
}

func extractGroupItem(r1 string, r2 string, r3 string) rune {
	itemsCount := map[rune]int{}

	addCount(itemsCount, getItemsCount(r1))
	addCount(itemsCount, getItemsCount(r2))
	addCount(itemsCount, getItemsCount(r3))
	for key, value := range itemsCount {
		if value == 3 {
			return key
		}
	}
	return -100
}

func addCount(out map[rune]int, in map[rune]bool) {
	for i, _ := range in {
		out[i]++
	}
}

func getItemsCount(ruck string) map[rune]bool {
	result := map[rune]bool{}
	for i := 0; i < len(ruck); i++ {
		result[rune(ruck[i])] = true
	}
	return result
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
