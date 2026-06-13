package bench

import (
	"math/rand"
	"testing"
)

const SIZE = 64 * 1024 * 1024

var globalSum int

func BenchmarkSequential(b *testing.B) {
	data := make([]int, SIZE)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0
		for i := 0; i < len(data); i++ {
			sum += data[i]
		}
		globalSum = sum
	}
}

func BenchmarkRandom(b *testing.B) {
	data := make([]int, SIZE)
	for i := range data {
		data[i] = i
	}
	indices := make([]int, SIZE)
	for i := range indices {
		indices[i] = rand.Intn(SIZE)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0
		for _, idx := range indices {
			sum += data[idx]
		}
		globalSum = sum
	}
}
