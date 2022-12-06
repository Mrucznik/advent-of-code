package main

import (
	_ "embed"
	"fmt"
)

// time needed 6:30

//go:embed input.txt
var input string

func main() {
	fmt.Println(findMarker(input))
}

func findMarker(msg string) int {
	for i := range msg {
		lettersInMarker := map[rune]struct{}{}
		for j := 0; j < 4; j++ {
			letter := rune(msg[i+j])
			if _, ok := lettersInMarker[letter]; ok {
				goto next
			} else {
				lettersInMarker[letter] = struct{}{}
			}
		}
		return i + 4
	next:
	}
	return -1
}
