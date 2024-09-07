package it_test

import (
	"iter"
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

// assertEqualSeq asserts that two sequences are equal.
func assertEqualSeq[V any](t *testing.T, expected, actual iter.Seq[V]) {
	t.Helper()
	expectedNext, expectedStop := iter.Pull(expected)
	defer expectedStop()
	actualNext, actualStop := iter.Pull(actual)
	defer actualStop()
	for {
		expectedV, expectedOK := expectedNext()
		actualV, actualOK := actualNext()
		assert.Equal(t, expectedOK, actualOK)
		assert.Equal(t, expectedV, actualV)
		if !expectedOK {
			return
		}
	}
}

// assertEqualSeq2 asserts that two sequences are equal.
func assertEqualSeq2[K comparable, V any](t *testing.T, expected, actual iter.Seq2[K, V]) {
	t.Helper()
	assert.Equal(t, maps.Collect(expected), maps.Collect(actual))
}
