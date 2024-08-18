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

func TestCount(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
	}{
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), want: 3},
		{name: "Empty", seq: it.All([]int{})},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Count(tc.seq)
			assert.Equal(t, tc.want, got)
		})
	}
}

func ExampleCount() {
	seq := it.All([]int{1, 2, 3})
	fmt.Println(it.Count(seq))
	// Output:
	// 3
}

func TestNth(t *testing.T) {
	seq := it.All([]int{1, 2, 3})
	testCases := []struct {
		name string
		n    int
		want int
		ok   bool
	}{
		{name: "Found", n: 1, want: 2, ok: true},
		{name: "NotFound", n: 3},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := it.Nth(seq, tc.n)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func ExampleNth() {
	seq := it.All([]int{1, 2, 3})
	fmt.Println(it.Nth(seq, 1))
	fmt.Println(it.Nth(seq, 3))
	// Output:
	// 2 true
	// 0 false
}

func TestLast(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
		ok   bool
	}{
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), want: 3, ok: true},
		{name: "Empty", seq: it.All([]int{})},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := it.Last(tc.seq)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func ExampleLast() {
	seq := it.All([]int{1, 2, 3})
	fmt.Println(it.Last(seq))
	seq = it.All([]int{})
	fmt.Println(it.Last(seq))
	// Output:
	// 3 true
	// 0 false
}