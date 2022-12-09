package main

import (
	"aoc/util"
	"math"
	"strings"
)

type Coord struct {
	x int
	y int
}

func partA(filename string) (int, error) {
	var knots = []Coord{{0, 0}, {0, 0}}
	return day9(filename, knots)
}

func day9(filename string, knots []Coord) (int, error) {
	file := util.MustReadFile(filename)

	var knotsVisitedMap = make(map[Coord]bool)

	for _, l := range strings.Split(string(file), "\n") {
		direction := l[0]
		steps := util.MustAtoI(string(l[2]))

		for i := 0; i < steps; i++ {
			switch direction {
			case 'R':
				knots[0].x += 1
			case 'L':
				knots[0].x -= 1
			case 'U':
				knots[0].y -= 1
			case 'D':
				knots[0].y += 1
			}

			for j := 1; j < len(knots); j++ {
				if math.Abs(float64(knots[j-1].x-knots[j].x)) > 1 ||
					math.Abs(float64(knots[j-1].y-knots[j].y)) > 1 {
					knots[j] = Coord{
						x: knots[j].x + orientation(knots[j].x, knots[j-1].x),
						y: knots[j].y + orientation(knots[j].y, knots[j-1].y),
					}
				}
			}
			knotsVisitedMap[knots[len(knots)-1]] = true
		}
	}

	return len(knotsVisitedMap), nil
}

func orientation(from int, to int) int {
	if to == from {
		return 0
	} else if to > from {
		return 1
	}
	return -1
}
