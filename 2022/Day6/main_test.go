package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_tuningTrouble_demo(t *testing.T) {
	index, err := demo()
	require.Nil(t, err)
	assert.Equal(t, 5, index)
}

func Test_tuningTrouble_part1(t *testing.T) {
	topStackCrates, err := part1()
	require.Nil(t, err)
	fmt.Println(topStackCrates)
}
