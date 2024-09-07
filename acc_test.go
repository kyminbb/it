package it_test

import (
	"cmp"
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
		{name: "Empty", seq: it.Empty[int]()},
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

	nums = it.Empty[int]()
	fmt.Println(it.Max(nums))
	// Output:
	// 3 true
	// 0 false
}

func TestMaxBy(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
		ok   bool
	}{
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), want: 1, ok: true},
		{name: "Empty", seq: it.Empty[int]()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := it.MaxBy(tc.seq, func(a, b int) int { return cmp.Compare(b, a) })
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func ExampleMaxBy() {
	type Student struct {
		name string
		age  int
	}
	compareStudents := func(a, b *Student) int {
		if c := cmp.Compare(a.age, b.age); c != 0 {
			return c
		}
		return cmp.Compare(a.name, b.name)
	}
	students := it.All([]*Student{{"Alice", 21}, {"Bob", 21}, {"Charlie", 20}})
	fmt.Println(it.MaxBy(students, compareStudents))

	students = it.Empty[*Student]()
	fmt.Println(it.MaxBy(students, compareStudents))
	// Output:
	// &{Bob 21} true
	// <nil> false
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
		ok   bool
	}{
		{name: "NonEmpty", seq: it.All([]int{3, 2, 1}), want: 1, ok: true},
		{name: "Empty", seq: it.Empty[int]()},
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

	nums = it.Empty[int]()
	fmt.Println(it.Min(nums))
	// Output:
	// 1 true
	// 0 false
}

func TestMinBy(t *testing.T) {
	testCases := []struct {
		name string
		seq  iter.Seq[int]
		want int
		ok   bool
	}{
		{name: "NonEmpty", seq: it.All([]int{1, 2, 3}), want: 3, ok: true},
		{name: "Empty", seq: it.Empty[int]()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := it.MinBy(tc.seq, func(a, b int) int { return cmp.Compare(b, a) })
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func ExampleMinBy() {
	type Student struct {
		name string
		age  int
	}
	compareStudents := func(a, b *Student) int {
		if c := cmp.Compare(a.age, b.age); c != 0 {
			return c
		}
		return cmp.Compare(a.name, b.name)
	}
	students := it.All([]*Student{{"Alice", 21}, {"Bob", 20}, {"Charlie", 20}})
	fmt.Println(it.MinBy(students, compareStudents))

	students = it.Empty[*Student]()
	fmt.Println(it.MinBy(students, compareStudents))
	// Output:
	// &{Bob 20} true
	// <nil> false
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
			seq:  it.Empty[int](),
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
