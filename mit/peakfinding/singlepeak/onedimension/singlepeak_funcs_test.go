package singlepeak_test

import (
	"testing"

	peak "github.com/shanehowearth/algorithms/mit/peakfinding/singlepeak/onedimension"
	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		result int
	}{
		"Single value input": {
			input:  []int{5},
			result: 0,
		},
		"All eq": {
			input:  []int{5, 5, 5, 5, 5, 5},
			result: 5,
		},
		"Single peak": {
			input:  []int{3, 5, 4, 2, 1},
			result: 1,
		},
		"End peak": {
			input:  []int{1, 2, 3, 4, 5},
			result: 4,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {})
		tp, err := peak.NewPeak(tc.input)
		assert.NotNil(t, tp, "No test peak.Peak created")
		assert.Nil(t, err, "Error generated creating peak.Peak")

		// Actual tests
		result := tp.Linear()
		assert.Equal(t, tc.result, result, "Got a different result")

	}
}

func TestDandC(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		result int
	}{
		"Single value input": {
			input:  []int{5},
			result: 0,
		},
		"All eq": {
			input:  []int{5, 5, 5, 5, 5, 5},
			result: 0, // Note: The DandC alg moves to the left when two vals are eq, but the Linear alg moves to the right
		},
		"Single peak": {
			input:  []int{3, 5, 4, 2, 1},
			result: 1,
		},
		"End peak": {
			input:  []int{1, 2, 3, 4, 5},
			result: 4,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {})
		tp, err := peak.NewPeak(tc.input)
		assert.NotNil(t, tp, "No test peak.Peak created")
		assert.Nil(t, err, "Error generated creating peak.Peak")

		// Actual tests
		result := tp.DandC()
		assert.Equal(t, tc.result, result, "Got a different result")

	}
}
