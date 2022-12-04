package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_campCleanup_demo(t *testing.T) {
	contained, overlaps, err := campCleanup("example-input.txt")
	require.Nil(t, err)
	assert.Equal(t, 2, contained)
	assert.Equal(t, 4, overlaps)
}

func Test_campCleanup(t *testing.T) {
	contained, overlaps, err := campCleanup("input.txt")
	require.Nil(t, err)
	fmt.Println(contained)
	fmt.Println(overlaps)
}
