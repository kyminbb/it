package it

import (
	"cmp"
	"iter"
)

// Max returns the maximum element of seq.
// The second return value reports whether the element exists (seq is not empty).
func Max[V cmp.Ordered](seq iter.Seq[V]) (V, bool) {
	var max V
	ok := false
	for v := range seq {
		ok = true
		if v > max {
			max = v
		}
	}
	return max, ok
}
