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

// Count iterates over seq and returns the number of elements.
func Count[V any](seq iter.Seq[V]) int {
	count := 0
	for range seq {
		count++
	}
	return count
}

// Nth returns the nth element of seq.
// The second return value reports whether the element exists.
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

// Last iterates over seq and returns the last element.
// The second return value reports whether the element exists (seq is not empty).
func Last[V any](seq iter.Seq[V]) (V, bool) {
	var v V
	ok := false
	for v = range seq {
		ok = true
	}
	return v, ok
}

// Cycle returns an iterator that repeates seq endlessly.
//
// If seq is empty, the returned iterator is also empty.
func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			next, stop := iter.Pull(seq)
			defer stop()
			isFirst := true
		once:
			for {
				v, ok := next()
				switch {
				case !ok:
					if isFirst {
						// seq is empty
						return
					}
					// Repeat seq
					break once
				case !yield(v):
					return
				}
				isFirst = false
			}
		}
	}
}
