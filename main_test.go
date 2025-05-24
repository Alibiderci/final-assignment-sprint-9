package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {

	t.Run("SizeOfZero", func(t *testing.T) {
		slice := generateRandomElements(0)
		assert.Nil(t, slice)
	})

	t.Run("NegativeSize", func(t *testing.T) {
		slice := generateRandomElements(-5)
		assert.Nil(t, slice)
	})

	size := 10
	slice := generateRandomElements(size)
	actualLength:= len(slice)
	expectedLength := size

	require.NotNil(t, slice)
	assert.Equal(t, expectedLength, actualLength)

}

func TestMaximum(t *testing.T) {
	testSlice1 := []int{15, 2, 24, 30, 1, 8, 59, 100, 777, 500}
	expected := 777

	actual := maximum(testSlice1)
	assert.Equal(t, expected, actual) 

	testSlice2 := []int{8, 59, 100, 340, 227}
	expected = 340
	actual = maximum(testSlice2)

	assert.Equal(t, expected, actual)
}

