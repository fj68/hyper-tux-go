package hyper

import (
	"math/rand"

	"github.com/fj68/hyper-tux-go/internal/slicetools"
)

// Choice performs weighted random selection from candidates using the given weights.
func Choice[T any](candidates []T, weights []int) T {
	total := slicetools.Sum(weights)
	r := rand.Intn(total)

	for i, w := range weights {
		r -= w
		if r < 0 {
			return candidates[i]
		}
	}
	var zero T
	return zero
}
