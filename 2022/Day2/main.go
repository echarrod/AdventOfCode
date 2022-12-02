package main

import (
	"os"
	"strings"
)

func rockPaperScissors() (int, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}

	totalPoints := 0

	rounds := strings.Split(string(file), "\n")
	for _, r := range rounds {
		roundSlice := strings.Fields(r)
		elfChoice := roundSlice[0]
		playerChoice := roundSlice[1]

		totalPoints += getClashPoints(playerChoice, elfChoice)
	}

	return totalPoints, nil
}

func getClashPoints(playerChoice string, elfChoice string) int {
	roundScore := 0
	switch playerChoice {
	case "X": // Rock
		roundScore += 1
		if elfChoice == "A" { // Rock
			roundScore += 3
		} else if elfChoice == "C" { // Scissors
			roundScore += 6
		} else {
			roundScore += 0
		}
	case "Y": // Paper
		roundScore += 2
		if elfChoice == "B" { // Paper
			roundScore += 3
		} else if elfChoice == "A" { // Rock
			roundScore += 6
		} else {
			roundScore += 0
		}
	case "Z": // Scissors
		roundScore += 3
		if elfChoice == "C" { // Scissors
			roundScore += 3
		} else if elfChoice == "B" { // Paper
			roundScore += 6
		} else {
			roundScore += 0
		}

	default:
		panic("Invalid player choice")
	}

	return roundScore
}
