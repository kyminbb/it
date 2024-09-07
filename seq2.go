package it

import "iter"

// All2 returns an iterator over all key-value pairs of m.
//
// The returned iterator yields the pairs in an arbitrary order.
func All2[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Zip returns an iterator over pairs of elements from seqK and seqV.
//
// The returned iterator stops when either seqK or seqV stops.
func Zip[K, V any](seqK iter.Seq[K], seqV iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		nextK, stopK := iter.Pull(seqK)
		defer stopK()
		nextV, stopV := iter.Pull(seqV)
		defer stopV()
		for {
			v, okV := nextK()
			w, okW := nextV()
			if !okV || !okW || !yield(v, w) {
				return
			}
		}
	}
}

// Unzip consumes all elements of seq and returns two iterators: one over the left elements and
// one over the right elements.
func Unzip[K, V any](seq iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	var ks []K
	var vs []V
	seq(func(k K, v V) bool {
		ks = append(ks, k)
		vs = append(vs, v)
		return true
	})
	return All(ks), All(vs)
}
