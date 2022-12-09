package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_demo(t *testing.T) {
	a, err := partA("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, 13, a)
}

func Test_input(t *testing.T) {
	a, err := partA("input.txt")
	require.Nil(t, err)
	fmt.Println(a)
}
