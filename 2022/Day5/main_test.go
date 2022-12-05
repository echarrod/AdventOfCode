package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_getTopStackCrates_demo(t *testing.T) {
	topStackCrates, err := parseCrates("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, "CMZ", topStackCrates)
}

func Test_getTopStackCrates(t *testing.T) {
	topStackCrates, err := parseCrates("input.txt")
	require.Nil(t, err)
	fmt.Println(topStackCrates)
}
