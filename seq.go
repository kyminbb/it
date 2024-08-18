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

// Empty returns an iterator that yields nothing.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {
	}
}

// Take returns an iterator that yields the first n elements of seq, or fewer if seq ends sooner.
func Take[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for i := 0; i < n; i++ {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

// Skip returns an iterator that skips the first n elements of seq and yields the rest.
func Skip[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for i := 0; i < n; i++ {
			if _, ok := next(); !ok {
				return
			}
		}
		for {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

// Cycle returns an iterator that repeats seq endlessly.
//
// If seq is empty, the returned iterator is also empty.
func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		// Reports whether the iterator should continue
		iterateOnce := func() bool {
			next, stop := iter.Pull(seq)
			defer stop()
			isFirst := true
			for {
				v, ok := next()
				switch {
				case !ok:
					if isFirst {
						// seq is empty
						return false
					}
					// Repeat seq
					return true
				case !yield(v):
					return false
				}
				isFirst = false
			}
		}
		for iterateOnce() {
		}
	}
}

// Repeat returns an iterator that yields v endlessly.
func Repeat[V any](v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for yield(v) {
		}
	}
}

// Chain returns an iterator that yields elements from each of the input iterators in turn.
func Chain[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		// Reports whether the iterator should continue
		iterate := func(seq iter.Seq[V]) bool {
			next, stop := iter.Pull(seq)
			defer stop()
			for {
				v, ok := next()
				switch {
				case !ok:
					return true
				case !yield(v):
					return false
				}
			}
		}
		for _, seq := range seqs {
			if !iterate(seq) {
				return
			}
		}
	}
}
