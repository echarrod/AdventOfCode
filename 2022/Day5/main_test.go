package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_getTopStackCrates_demo(t *testing.T) {
	topStackCrates, err := CrateMover9000("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, "CMZ", topStackCrates)
}

func Test_getTopStackCrates_a(t *testing.T) {
	topStackCrates, err := CrateMover9000("input.txt")
	require.Nil(t, err)
	fmt.Println(topStackCrates)
}

func Test_getTopStackCrates_b(t *testing.T) {
	topStackCrates, err := CrateMover9001("input.txt")
	require.Nil(t, err)
	fmt.Println(topStackCrates)
}
