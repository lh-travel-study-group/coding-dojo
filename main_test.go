package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallenge(t *testing.T) {

	alice := []int{30, 0, 40}
	bob := []int{40, 0, 30}

	result, err := CalcScores(alice, bob)

	assert.Equal(t, []int{1, 1}, result)
	assert.Empty(t, err)
}

func TestValidation(t *testing.T) {

	alice := []int{30, 0, 40, 5}
	bob := []int{40, 0, 30}

	result, err := CalcScores(alice, bob)

	assert.Nil(t, result)
	assert.Equal(t, "all slices must have 3 entries", err)

}
