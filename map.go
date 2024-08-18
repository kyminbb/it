package it

import "iter"

// Map returns an iterator that calls f on each element of seq.
func Map[V, R any](seq iter.Seq[V], f func(V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
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
