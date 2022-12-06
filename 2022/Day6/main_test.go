package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_tuningTrouble_demo(t *testing.T) {
	index, _, err := indexOfStart("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, 5, index)
}

func Test_tuningTrouble(t *testing.T) {
	part1, part2, err := indexOfStart("input.txt")
	require.Nil(t, err)
	fmt.Println(part1)
	fmt.Println(part2)
}
