package main

import (
	"errors"
	"os"
)

func demo() (int, error) {
	return part("example-input.txt")
}

func part1() (int, error) {
	return part("input.txt")
}

func part(input string) (int, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return 0, err
	}

	line := string(file)
	for i := 3; i < len(line); i++ {
		if unique(line[i-3 : i+1]) {
			return i + 1, nil
		}
	}

	return 0, errors.New("unique chars not found")
}

func unique(s string) bool {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	return len(set) == len(s)
}
