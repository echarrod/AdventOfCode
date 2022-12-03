package main

import (
	"os"
	"strings"
)

func rucksackReorganisation(filename string) (int, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	sumPriorities := 0

	rucksacks := strings.Split(string(file), "\n")
	for _, r := range rucksacks {
		mid := len(r) / 2
		leftPart := r[:mid]
		rightPart := r[mid:]

		for _, lc := range leftPart {
			if strings.ContainsRune(rightPart, lc) {
				sumPriorities += getPriority(lc)
				break
			}
		}
	}

	return sumPriorities, nil
}

func getPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}
