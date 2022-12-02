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

var losingMoves = map[rune]rune{
	'B': 'A',
	'C': 'B',
	'A': 'C',
}

func main() {
	rows := strings.Split(input, "\n")
	mySum, oppSum := calculateStrategyOutcome(rows)
	fmt.Printf("My points: %d\n", mySum)
	fmt.Printf("Opp poitns %d\n", oppSum)
}

func calculateStrategyOutcome(rows []string) (int, int) {
	mySum, oppSum := 0, 0
	for _, row := range rows {
		myPoints, oppPoints := calculatePoints(rune(row[0]), rune(row[2]))
		mySum += myPoints
		oppSum += oppPoints
	}
	return mySum, oppSum
}

func calculatePoints(oppPlay, whatMove rune) (myPoints, oppPoints int) {
	myPlay := convertMoveToPlay(whatMove, oppPlay)

	myPoints, oppPoints = getRoundPoints(myPlay, oppPlay)

	myPoints += getPlayPoints(myPlay)
	oppPoints += getPlayPoints(oppPlay)
	return
}

func convertMoveToPlay(move, oppPlay rune) rune {
	if move == 'X' {
		// lose
		return losingMoves[oppPlay]
	} else if move == 'Y' {
		// draw
		return oppPlay
	} else {
		// win
		return winningMoves[oppPlay]
	}
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

func getPlayPoints(play rune) int {
	switch play {
	case 'A':
		return 1
	case 'B':
		return 2
	case 'C':
		return 3
	}
	return -100000000000
}
