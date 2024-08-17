package it

import "iter"

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
