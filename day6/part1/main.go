package main

import (
	_ "embed"
	"fmt"
)

// time needed 1min 43s

//go:embed input.txt
var input string

func main() {
	fmt.Println(findMarker(input, 14))
}

func findMarker(msg string, markLength int) int {
	for i := range msg {
		lettersInMarker := map[rune]struct{}{}
		for j := 0; j < markLength; j++ {
			letter := rune(msg[i+j])
			if _, ok := lettersInMarker[letter]; ok {
				goto next
			} else {
				lettersInMarker[letter] = struct{}{}
			}
		}
		return i + markLength
	next:
	}
	return -1
}
