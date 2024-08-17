package it

import "iter"

// All returns an iterator over all elements of s.
func All[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// Count consumes seq and returns the number of elements.
func Count[V any](seq iter.Seq[V]) int {
	count := 0
	for range seq {
		count++
	}
	return count
}

// Nth returns the n-th element of seq. The second return value reports whether the element exists.
//
// The count is zero-based, so Nth(seq, 0) returns the first element.
func Nth[V any](seq iter.Seq[V], n int) (V, bool) {
	i := 0
	for v := range seq {
		if i == n {
			return v, true
		}
		i++
	}
	var zero V
	return zero, false
}
