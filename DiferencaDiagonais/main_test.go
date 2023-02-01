package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var matriz = [][]int{{1, 2, 3},
	{4, 5, 6},
	{9, 8, 9}}

func TestDiagonal(t *testing.T) {
	result := GetDiagonalDiff(matriz)
	assert.Equal(t, 2, result)
}

func TestPrimaryDiagonal(t *testing.T) {
	result := getPrimaryDiagonalSum(matriz)
	assert.Equal(t, 15, result)
}

func TestSecondaryDiagonal(t *testing.T) {
	result := getSecondaryDiagonalSum(matriz)
	assert.Equal(t, 17, result)
}
