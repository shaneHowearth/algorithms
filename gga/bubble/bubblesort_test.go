package bubble_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/shanehowearth/algorithms/gga/bubble"
)

func TestBubble(t *testing.T) {
	testcases := map[string]struct {
		before []int
		after  []int
	}{
		"Book example": {
			before: []int{90, 70, 50, 80, 60, 85},
			after:  []int{50, 60, 70, 80, 85, 90},
		},
		"Zero length": {
			before: []int{},
			after:  []int{},
		},
		"Already sorted": {
			before: []int{50, 60, 70, 80, 85, 90},
			after:  []int{50, 60, 70, 80, 85, 90},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			bubble.Sort(tc.before)
			assert.Equal(t, len(tc.after), len(tc.before), "Array lengths do not match")
			for idx := range tc.after {
				assert.Equalf(t, tc.after[idx], tc.before[idx], "Array indicies at %d did not match", idx)
			}
		})
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("rand seeded")
}
func BenchmarkSort(b *testing.B) {
	input := rand.Perm(b.N)
	for i := 0; i < b.N; i++ {
		bubble.Sort(input[:i])
	}
}
