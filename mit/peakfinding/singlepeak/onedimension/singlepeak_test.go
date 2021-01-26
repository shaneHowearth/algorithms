package singlepeak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testcases := map[string]struct {
		input  []int
		output *Peak
		err    error
	}{
		"No input": {
			err: fmt.Errorf("no input to check"),
		},
		"Single element input": {
			input:  []int{1},
			output: &Peak{input: []int{1}},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			output, err := NewPeak(tc.input)
			if tc.output == nil {
				assert.Nil(t, output, "Wasn't expecting a concrete instance")
			} else {
				assert.NotNil(t, output, "Was expecting a concrete instance of peak.Peak")
				assert.Equal(t, tc.output, output, "Was not an instance of peak.Peak")
			}
			if tc.err == nil {
				assert.Nil(t, err, "Wasn't expecting a concrete instance")
			} else {
				assert.NotNil(t, err, "Was expecting a concrete instance of peak.Peak")
				assert.EqualError(t, tc.err, err.Error(), "Was not an instance of peak.Peak")
			}
		})
	}
}
