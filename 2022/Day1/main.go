package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxCalories() (int, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}

	splitCalories := strings.Split(string(file), "\n")
	var elfCalorieTotals []int
	currentElf := 0

	for _, c := range splitCalories {
		if c != "" {
			i, err := strconv.Atoi(c)
			if err != nil {
				return 0, err
			}
			
			currentElf += i
		} else {
			elfCalorieTotals = append(elfCalorieTotals, currentElf)
			currentElf = 0
		}
	}

	top3Calories := 0
	sort.Ints(elfCalorieTotals)
	for _, i := range elfCalorieTotals[len(elfCalorieTotals)-3:] {
		top3Calories += i
	}

	return top3Calories, nil
}
