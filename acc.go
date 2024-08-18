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
		if !ok || v > max {
			max, ok = v, true
		}
	}
	return max, ok
}

// Min returns the minimum element of seq.
// The second return value reports whether the element exists (seq is not empty).
func Min[V cmp.Ordered](seq iter.Seq[V]) (V, bool) {
	var min V
	ok := false
	for v := range seq {
		if !ok || v < min {
			min, ok = v, true
		}
	}
	return min, ok
}
