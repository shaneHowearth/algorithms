package builtin

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := rand.Perm(b.N)
		b.StartTimer()
		sort.Ints(input)
	}
}

func BenchmarkCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := rand.Perm(b.N)
		b.StartTimer()
		sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	}
}
