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

func getClashPoints(desiredOutcome string, elfChoice string) int {
	roundScore := 0
	switch elfChoice {
	case "A": // Rock
		if desiredOutcome == "X" { // Lose
			roundScore += 3 // 3 for scissors, 0 for losing
			break
		} else if desiredOutcome == "Y" { // Draw
			roundScore += 4 // 1 for rock, 3 for drawing
			break
		} else if desiredOutcome == "Z" { // Win
			roundScore += 8 // 2 for paper, 6 for winning
			break
		}
		panic("Invalid outcome")

	case "B": // Paper
		if desiredOutcome == "X" { // Lose
			roundScore += 1 // 1 for rock, 0 for losing
			break
		} else if desiredOutcome == "Y" { // Draw
			roundScore += 5 // 2 for paper, 3 for drawing
			break
		} else if desiredOutcome == "Z" { // Win
			roundScore += 9 // 3 for scissors, 6 for winning
			break
		}
		panic("Invalid outcome")

	case "C": // Scissors
		if desiredOutcome == "X" { // Lose
			roundScore += 2 // 2 for paper, 0 for losing
			break
		} else if desiredOutcome == "Y" { // Draw
			roundScore += 6 // 3 for scissors, 3 for drawing
			break
		} else if desiredOutcome == "Z" { // Win
			roundScore += 7 // 1 for rock, 6 for winning
			break
		}
		panic("Invalid outcome")

	default:
		panic("Invalid player choice")
	}

	return roundScore
}
