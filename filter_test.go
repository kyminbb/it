package it_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/kyminbb/it"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name      string
		seq       iter.Seq[int]
		predicate func(int) bool
		want      iter.Seq[int]
	}{
		{
			name:      "Positive",
			seq:       it.All([]int{1, -2, 3}),
			predicate: func(x int) bool { return x > 0 },
			want:      it.All([]int{1, 3}),
		},
		{
			name:      "Empty",
			seq:       it.All([]int{}),
			predicate: func(x int) bool { return true },
			want:      it.All([]int{}),
		},
		{
			name:      "All",
			seq:       it.All([]int{1, 2, 3}),
			predicate: func(x int) bool { return true },
			want:      it.All([]int{1, 2, 3}),
		},
		{
			name:      "None",
			seq:       it.All([]int{1, 2, 3}),
			predicate: func(x int) bool { return false },
			want:      it.All([]int{}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Filter(tc.seq, tc.predicate)
			assertEqualSeq(t, tc.want, got)
		})
	}
}

func ExampleFilter() {
	nums := it.All([]int{1, -2, 3})
	positives := it.Filter(nums, func(x int) bool { return x > 0 })
	for v := range positives {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 3
}
