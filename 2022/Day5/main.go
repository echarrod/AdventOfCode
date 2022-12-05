package main

import (
	"aoc/util"
	"os"
	"strings"
)

type Movement struct {
	move int
	from int
	to   int
}

func parseCrates(filename string) (topOfStacks string, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	maxCrate := 0
	readyForInstructions := false
	stacks := map[int]string{}
	var moves []Movement

	for _, line := range strings.Split(string(file), "\n") {
		if line == "" || line[1] == '1' {
			readyForInstructions = true
			continue
		}

		if readyForInstructions {
			lineFields := strings.Fields(line)
			movement := Movement{
				move: util.MustAtoI(lineFields[1]),
				from: util.MustAtoI(lineFields[3]),
				to:   util.MustAtoI(lineFields[5]),
			}
			moves = append(moves, movement)
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(line), "[") {
			// 4 chars between each create letter
			for i := 0; i < len(line); i += 4 {
				crateNo := (i / 4) + 1
				if crateNo > maxCrate {
					maxCrate = crateNo
				}

				// if crate present
				if line[i+1] != ' ' {
					stacks[crateNo] += string(line[i+1])
				}
			}
		}

	}

	for _, move := range moves {
		for i := 0; i < move.move; i++ {
			stacks[move.to] = string(stacks[move.from][0]) + stacks[move.to]
			stacks[move.from] = stacks[move.from][1:]
		}
	}

	var result strings.Builder
	for i := 1; i <= maxCrate; i++ {
		result.WriteRune(rune(stacks[i][0]))
	}

	return result.String(), nil
}
