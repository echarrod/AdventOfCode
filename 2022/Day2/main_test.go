package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rockPaperScissors(t *testing.T) {
	score, err := rockPaperScissors()
	require.Nil(t, err)
	fmt.Println(score)
}
