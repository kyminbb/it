package it_test

import (
	"iter"
	"testing"

	"github.com/kyminbb/it"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		f    func(int) int
		want iter.Seq[int]
	}{
		{
			name: "Double",
			seq:  it.All([]int{1, 2, 3}),
			f:    func(x int) int { return x * 2 },
			want: it.All([]int{2, 4, 6}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Map(tc.seq, tc.f)
			assertEqualSeq(t, tc.want, got)
		})
	}
}
