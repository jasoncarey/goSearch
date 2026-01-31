package engine

import (
	"fmt"
	"math/rand"
	"testing"

	"search/vec"
)

func RandomVector(dim int) vec.Vector {
	v := make(vec.Vector, dim)
	for i := range dim {
		v[i] = rand.NormFloat64()
	}

	return v
}

func BenchmarkBruteForceSearch(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			e := &Engine{}
			for range size {
				e.Add(RandomVector(128))
			}
			query := RandomVector(128)
			b.ResetTimer()

			for b.Loop() {
				e.BruteForceSearch(query, 10)
			}
		})
	}
}
