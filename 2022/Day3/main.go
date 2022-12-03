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
	for k, _ := range rucksacks {
		if k > 0 && (k+1)%3 == 0 {
			for _, b := range rucksacks[k] {
				if strings.ContainsRune(rucksacks[k-1], b) && strings.ContainsRune(rucksacks[k-2], b) {
					sumPriorities += getPriority(b)
					break
				}
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
