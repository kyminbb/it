package it_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/kyminbb/it"
	"github.com/stretchr/testify/assert"
)

func ExampleAll() {
	seq := it.All([]int{1, 2, 3})
	for v := range seq {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}

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
