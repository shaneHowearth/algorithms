package unorderedlinear_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	unorderedlinear "github.com/shanehowearth/algorithms/search/unordered_linear"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testcases := map[string]struct {
		input []int
		find  int
		err   error
		found int
	}{
		"Nothing to search": {
			err: fmt.Errorf("nothing to search"),
		},
		"Nothing found": {
			input: []int{1, 2, 3, 78, 5},
			find:  6,
			err:   fmt.Errorf("could not find 6"),
		},
		"First": {
			input: []int{6, 2, 3, 78, 5},
			found: 0,
			find:  6,
		},
		"Last": {
			input: []int{6, 2, 3, 78, 5},
			found: 4,
			find:  5,
		},
		"Mid(ish)": {
			input: []int{6, 2, 3, 78, 5},
			found: 2,
			find:  3,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			found, err := unorderedlinear.Search(tc.find, tc.input)
			if tc.err != nil {
				assert.NotNilf(t, err, "got nil when expecting an error that matches %v", tc.err)
				assert.EqualError(t, tc.err, err.Error(), "errors don't match")
			} else {
				assert.Nilf(t, err, "expected nil, but got an error %v", err)
				assert.Equal(t, tc.found, found, "got different positions back")
			}
		})
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func BenchmarkSearch(b *testing.B) {
	input := rand.Perm(b.N)
	find := rand.Intn(b.N)
	for i := 0; i < b.N; i++ {
		unorderedlinear.Search(find, input)
	}
}
