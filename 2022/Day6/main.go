package main

import (
	"os"
)

func indexOfStart(input string) (partA int, partB int, err error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return 0, 0, err
	}

	line := string(file)
	for i := 3; i < len(line); i++ {
		if isUnique(line[i-3:i+1]) && partA == 0 {
			partA = i + 1
		}

		if i > 12 && isUnique(line[i-13:i+1]) && partB == 0 {
			partB = i + 1
		}
	}

	return partA, partB, nil
}

func isUnique(s string) bool {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	return len(set) == len(s)
}
