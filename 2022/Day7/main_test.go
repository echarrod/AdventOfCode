package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func Test_directory_demo(t *testing.T) {
	size, err := sumDirectorySizes("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, 95437, size)
}

func Test_day6(t *testing.T) {
	size, err := sumDirectorySizes("input.txt")
	require.Nil(t, err)
	fmt.Println(size)
}
