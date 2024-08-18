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
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), want: 3, ok: true},
		{name: "Empty", seq: it.All([]int{})},
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
	nums := it.All([]int{1, 2, 3})
	fmt.Println(it.Max(nums))

	nums = it.All([]int{})
	fmt.Println(it.Max(nums))
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
		{name: "NonEmpty", seq: it.All([]int{3, 2, 1}), want: 1, ok: true},
		{name: "Empty", seq: it.All([]int{})},
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
	nums := it.All([]int{3, 2, 1})
	fmt.Println(it.Min(nums))

	nums = it.All([]int{})
	fmt.Println(it.Min(nums))
	// Output:
	// 1 true
	// 0 false
}

func TestFold(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		init int
		f    func(int, int) int
		want int
	}{
		{
			name: "Sum",
			seq:  it.All([]int{1, 2, 3}),
			f:    func(acc, v int) int { return acc + v },
			want: 6,
		},
		{
			name: "Product",
			seq:  it.All([]int{3, 4, 5}),
			init: 1,
			f:    func(acc, v int) int { return acc * v },
			want: 60,
		},
		{
			name: "Empty",
			seq:  it.All([]int{}),
			init: 1,
			f:    func(acc, v int) int { return acc * v },
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Fold(tc.seq, tc.init, tc.f)
			assert.Equal(t, tc.want, got)
		})
	}
}

func ExampleFold() {
	nums := it.All([]int{1, 2, 3})
	sum := it.Fold(nums, 0, func(acc, v int) int { return acc + v })
	fmt.Println(sum)

	nums = it.All([]int{3, 4, 5})
	product := it.Fold(nums, 1, func(acc, v int) int { return acc * v })
	fmt.Println(product)
	// Output:
	// 6
	// 60
}
