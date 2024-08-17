package it

import "iter"

// Filter returns an iterator that yields only the elements of seq for which predicate is true.
func Filter[V any](seq iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok || (predicate(v) && !yield(v)) {
				return
			}
		}
	}
}
