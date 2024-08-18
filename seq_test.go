package it_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/kyminbb/it"
	"github.com/stretchr/testify/assert"
)

func ExampleAll() {
	nums := it.All([]int{1, 2, 3})
	for v := range nums {
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
	nums := it.All([]int{1, 2, 3})
	fmt.Println(it.Count(nums))
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
	nums := it.All([]int{1, 2, 3})
	fmt.Println(it.Nth(nums, 1))
	fmt.Println(it.Nth(nums, 3))
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
	nums := it.All([]int{1, 2, 3})
	fmt.Println(it.Last(nums))

	nums = it.All([]int{})
	fmt.Println(it.Last(nums))
	// Output:
	// 3 true
	// 0 false
}

func TestTake(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		n    int
		want iter.Seq[int]
	}{
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), n: 2, want: it.All([]int{1, 2})},
		{name: "EarlyEnd", seq: it.All([]int{1, 2, 3}), n: 4, want: it.All([]int{1, 2, 3})},
		{name: "Empty", seq: it.All([]int{}), n: 2, want: it.All([]int{})},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Take(tc.seq, tc.n)
			assertEqualSeq(t, tc.want, got)
		})
	}
}

func ExampleTake() {
	nums := it.All([]int{1, 2, 3})
	take := it.Take(nums, 2)
	for v := range take {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
}

func TestSkip(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		n    int
		want iter.Seq[int]
	}{
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), n: 2, want: it.All([]int{3})},
		{name: "EarlyEnd", seq: it.All([]int{1, 2, 3}), n: 4, want: it.All([]int{})},
		{name: "Empty", seq: it.All([]int{}), n: 2, want: it.All([]int{})},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Skip(tc.seq, tc.n)
			assertEqualSeq(t, tc.want, got)
		})
	}
}

func ExampleSkip() {
	nums := it.All([]int{1, 2, 3})
	skip := it.Skip(nums, 2)
	for v := range skip {
		fmt.Println(v)
	}
	// Output:
	// 3
}

func TestCycle(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
	}{
		{name: "NonEmpty", s: []int{1, 2, 3}},
		{name: "Empty", s: []int{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Cycle(it.All(tc.s))
			i, count := 0, len(tc.s)
			for v := range got {
				if i == count*3 {
					break
				}
				assert.Equal(t, tc.s[i%count], v)
				i++
			}
		})
	}
}

func ExampleCycle() {
	nums := it.All([]int{1, 2, 3})
	cycle := it.Cycle(nums)
	i := 0
	for v := range cycle {
		if i == 6 {
			break
		}
		fmt.Println(v)
		i++
	}
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestRepeat(t *testing.T) {
	got := it.Repeat(3)
	i := 0
	for v := range got {
		if i == 3 {
			break
		}
		assert.Equal(t, 3, v)
		i++
	}
}

func ExampleRepeat() {
	repeat := it.Repeat(3)
	i := 0
	for v := range repeat {
		if i == 3 {
			break
		}
		fmt.Println(v)
		i++
	}
	// Output:
	// 3
	// 3
	// 3
}
