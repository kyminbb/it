package it

import "iter"

// All2 returns an iterator over all key-value pairs of m.
func All2[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Zip returns an iterator over pairs of elements from seqV and seqW.
//
// The returned iterator stops when either seqV or seqW stops.
func Zip[V, W any](seqV iter.Seq[V], seqW iter.Seq[W]) iter.Seq2[V, W] {
	return func(yield func(V, W) bool) {
		nextV, stopV := iter.Pull(seqV)
		defer stopV()
		nextW, stopW := iter.Pull(seqW)
		defer stopW()
		for {
			v, okV := nextV()
			w, okW := nextW()
			if !okV || !okW || !yield(v, w) {
				return
			}
		}
	}
}
