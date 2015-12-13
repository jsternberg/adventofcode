package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert := assert.New(t)
	paper, ribbon := calculateWrappingPaper(2, 3, 4)
	assert.Equal(58, paper)
	assert.Equal(34, ribbon)
	paper, ribbon = calculateWrappingPaper(1, 1, 10)
	assert.Equal(43, paper)
	assert.Equal(14, ribbon)
}

func TestParseLine(t *testing.T) {
	assert := assert.New(t)
	l, w, h, err := parseLine("3x4x5")
	if assert.NoError(err) {
		assert.Equal(3, l)
		assert.Equal(4, w)
		assert.Equal(5, h)
	}
}
