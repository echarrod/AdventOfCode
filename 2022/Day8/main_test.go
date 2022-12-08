package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_demo(t *testing.T) {
	a, _, err := treetopTreeHouse("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, 21, a)
}

func Test_input(t *testing.T) {
	a, b, err := treetopTreeHouse("input.txt")
	require.Nil(t, err)
	fmt.Println(a)
	fmt.Println(b)
}
