package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		name string
		inputSize int
		expectNil bool 
		expectedLength int
		checkElementsRange bool
	}{
		{
			name: "SizeOfZero",
			inputSize: 0,
			expectNil: true,
		},
		{
			name: "NegativeSize",
			inputSize: -5,
			expectNil: true,
		},
		{
			name: "PositiveSizeOne",
			inputSize: 1,
			expectNil: false,
			expectedLength: 1,
			checkElementsRange: true,
		},
		{
			name: "PositiveSmallSize",
			inputSize: 100,
			expectNil: false,
			expectedLength: 100,
			checkElementsRange: true,
		},
		{
			name: "PositiveBigSize",
			inputSize: 100_000,
			expectNil: false,
			expectedLength: 100_000,
			checkElementsRange: true,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := generateRandomElements(tc.inputSize)

			if tc.expectNil {
				assert.Nil(t, actual)
			} else {
				require.NotNil(t, actual)
				assert.Equal(t, tc.expectedLength, len(actual))
				if tc.checkElementsRange {
					for _, val := range actual {
						assert.True(t, val >= 1 && val <= 1000)
					}
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		expected int
	}{
		{
			name: "NilCase",
			input: nil,
			expected: 0, 
		},
		{
			name: "EmptySlice",
			input: []int{},
			expected: 0,
		},
		{
			name: "SingleElement",
			input: []int{5},
			expected: 5,
		},
		{
			name: "NegativeNumbers",
			input: []int{-5, -2, -10, -1},
			expected: -1,
		},
		{
			name: "PositiveNumbers",
			input: []int{5, 2, 10, 1},
			expected: 10,
		},
		{
			name: "MixedPositiveAndNegative",
			input: []int{5, -2, -10, 1},
			expected: 5,
		},
		{
			name: "AllSameNumber",
			input: []int{7, 7, 7, 7},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maximum(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

