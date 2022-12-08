package main

import (
	"aoc/util"
	"strings"
)

func treetopTreeHouse(filename string) (int, int, error) {
	file := util.MustReadFile(filename)
	highestScenicScore := 0

	var lines []string
	for _, l := range strings.Split(string(file), "\n") {
		if l != "" {
			lines = append(lines, l)
		}
	}

	// edge trees
	visibleCount := len(lines[0])*2 + len(lines)*2 - 4

	// start and finish 1 from the outside since we know these are visible
	for x := 1; x < len(lines[0])-1; x++ {
		for y := 1; y < len(lines)-1; y++ {
			i := 0
			score := 1
			isVisible := false

			// check for visibility to left
			for i = x - 1; i > 0 && lines[y][i] < lines[y][x]; i-- {
			}
			isVisible = i == 0 && lines[y][i] < lines[y][x]
			score = score * (x - i)

			// check for visibility to right
			for i = x + 1; i < len(lines[y])-1 && lines[y][i] < lines[y][x]; i++ {
			}
			isVisible = isVisible || i == len(lines[y])-1 && lines[y][i] < lines[y][x]
			score = score * (i - x)

			// check for visibility to top
			for i = y - 1; i > 0 && lines[i][x] < lines[y][x]; i-- {
			}
			isVisible = isVisible || i == 0 && lines[i][x] < lines[y][x]
			score = score * (y - i)

			// check for visibility to bottom
			for i = y + 1; i < len(lines)-1 && lines[i][x] < lines[y][x]; i++ {
			}
			isVisible = isVisible || i == len(lines)-1 && lines[i][x] < lines[y][x]
			score = score * (i - y)

			if score > highestScenicScore {
				highestScenicScore = score
			}
			if isVisible {
				visibleCount += 1
			}
		}
	}

	return visibleCount, highestScenicScore, nil
}
