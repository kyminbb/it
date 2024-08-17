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
