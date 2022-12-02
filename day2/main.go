package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var winningMoves = map[rune]rune{
	'A': 'B',
	'B': 'C',
	'C': 'A',
}

func main() {
	mySum, oppSum := 0, 0
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		myPoints, oppPoints := calculatePoints(normalisePlay(rune(row[2])), normalisePlay(rune(row[0])))
		mySum += myPoints
		oppSum += oppPoints
	}
	fmt.Printf("My points: %d\n", mySum)
	fmt.Printf("Opp poitns %d\n", oppSum)
}

func calculatePoints(myPlay, oppPlay rune) (myPoints, oppPoints int) {
	myPoints, oppPoints = getRoundPoints(myPlay, oppPlay)

	myPoints += getPlayPoints(myPlay)
	oppPoints += getPlayPoints(oppPlay)
	return
}

func getRoundPoints(myPlay, oppPlay rune) (myPoints, oppPoints int) {
	if myPlay == oppPlay {
		return 3, 3
	}

	// won moves
	if winningMoves[oppPlay] == myPlay {
		return 6, 0
	} else {
		return 0, 6
	}
}

func normalisePlay(play rune) rune {
	switch play {
	case 'X':
		return 'A'
	case 'Y':
		return 'B'
	case 'Z':
		return 'C'
	}
	return play
}

func getPlayPoints(play rune) int {
	switch play {
	case 'A', 'X':
		return 1
	case 'B', 'Y':
		return 2
	case 'C', 'Z':
		return 3
	}
	return -100000000000
}
