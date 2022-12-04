package main

import (
	"os"
	"strconv"
	"strings"
)

func campCleanup(filename string) (int, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	contained := 0
	for _, lines := range strings.Fields(string(file)) {
		assignments := strings.Split(lines, ",")

		elf1Range := strings.Split(assignments[0], "-")
		elf1Start, _ := strconv.Atoi(elf1Range[0])
		elf1Finish, _ := strconv.Atoi(elf1Range[1])

		elf2Range := strings.Split(assignments[1], "-")
		elf2Start, _ := strconv.Atoi(elf2Range[0])
		elf2Finish, _ := strconv.Atoi(elf2Range[1])

		if elf1Start >= elf2Start && elf1Finish <= elf2Finish ||
			elf2Start >= elf1Start && elf2Finish <= elf1Finish {
			contained++
		}
	}

	return contained, nil
}
