package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rucksackReorganisation_demo(t *testing.T) {
	score, err := rucksackReorganisation("example-input.txt")
	require.Nil(t, err)
	fmt.Println(score)
}

func Test_rucksackReorganisation(t *testing.T) {
	score, err := rucksackReorganisation("input.txt")
	require.Nil(t, err)
	fmt.Println(score)
}
