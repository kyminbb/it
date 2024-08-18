package it_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/kyminbb/it"
	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
		ok   bool
	}{
		{"Empty", it.All([]int{}), 0, false},
		{"NonEmpty", it.All([]int{1, 2, 3}), 3, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := it.Max(tc.seq)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func ExampleMax() {
	seq := it.All([]int{1, 2, 3})
	fmt.Println(it.Max(seq))
	seq = it.All([]int{})
	fmt.Println(it.Max(seq))
	// Output:
	// 3 true
	// 0 false
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
		ok   bool
	}{
		{"Empty", it.All([]int{}), 0, false},
		{"NonEmpty", it.All([]int{3, 2, 1}), 1, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := it.Min(tc.seq)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func ExampleMin() {
	seq := it.All([]int{3, 2, 1})
	fmt.Println(it.Min(seq))
	seq = it.All([]int{})
	fmt.Println(it.Min(seq))
	// Output:
	// 1 true
	// 0 false
}
