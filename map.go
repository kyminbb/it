package it

import "iter"

// All returns an iterator over all elements in s.
func All[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// Map returns an iterator which calls f on each element of seq.
func Map[V, F any](seq iter.Seq[V], f func(V) F) iter.Seq[F] {
	return func(yield func(F) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok || !yield(f(v)) {
				return
			}
		}
	}
}
