package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCalories(t *testing.T) {
	calories, err := maxCalories()
	require.Nil(t, err)
	fmt.Println(calories)
}
