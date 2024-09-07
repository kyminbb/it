package it_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/kyminbb/it"
)

func ExampleAll2() {
	pairs := it.All2(map[int]string{1: "one", 2: "two", 3: "three"})
	for k, v := range pairs {
		fmt.Println(k, v)
	}
	// Output:
	// 1 one
	// 2 two
	// 3 three
}

func TestZip(t *testing.T) {
	type testCase struct {
		name string
		seqK iter.Seq[int]
		seqV iter.Seq[string]
		want iter.Seq2[int, string]
	}
	testCases := []testCase{
		{
			name: "Equal",
			seqK: it.All([]int{1, 2, 3}),
			seqV: it.All([]string{"one", "two", "three"}),
			want: it.All2(map[int]string{1: "one", 2: "two", 3: "three"}),
		},
		{
			name: "KShorter",
			seqK: it.All([]int{1, 2}),
			seqV: it.All([]string{"one", "two", "three"}),
			want: it.All2(map[int]string{1: "one", 2: "two"}),
		},
		{
			name: "VShorter",
			seqK: it.All([]int{1, 2, 3}),
			seqV: it.All([]string{"one", "two"}),
			want: it.All2(map[int]string{1: "one", 2: "two"}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := it.Zip(tc.seqK, tc.seqV)
			assertEqualSeq2(t, tc.want, got)
		})
	}
}

func ExampleZip() {
	nums := it.All([]int{1, 2, 3})
	words := it.All([]string{"one", "two", "three"})
	pairs := it.Zip(nums, words)
	for k, v := range pairs {
		fmt.Println(k, v)
	}
	fmt.Println()

	shortNums := it.All([]int{1, 2})
	pairs = it.Zip(shortNums, words)
	for k, v := range pairs {
		fmt.Println(k, v)
	}
	fmt.Println()

	shortWords := it.All([]string{"one", "two"})
	pairs = it.Zip(nums, shortWords)
	for k, v := range pairs {
		fmt.Println(k, v)
	}
	// Output:
	// 1 one
	// 2 two
	// 3 three
	//
	// 1 one
	// 2 two
	//
	// 1 one
	// 2 two
}

func TestUnzip(t *testing.T) {
	type testCase struct {
		name  string
		seq   iter.Seq2[int, string]
		wantK iter.Seq[int]
		wantV iter.Seq[string]
	}
	testCases := []testCase{
		{
			name:  "NonEmpty",
			seq:   it.All2(map[int]string{1: "one", 2: "two", 3: "three"}),
			wantK: it.All([]int{1, 2, 3}),
			wantV: it.All([]string{"one", "two", "three"}),
		},
		{
			name:  "Empty",
			seq:   it.Empty2[int, string](),
			wantK: it.Empty[int](),
			wantV: it.Empty[string](),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotK, gotV := it.Unzip(tc.seq)
			assertEqualSeq(t, tc.wantK, gotK)
			assertEqualSeq(t, tc.wantV, gotV)
		})
	}
}

func ExampleUnzip() {
	pairs := it.All2(map[int]string{1: "one", 2: "two", 3: "three"})
	nums, words := it.Unzip(pairs)
	for k := range nums {
		fmt.Println(k)
	}
	for v := range words {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// one
	// two
	// three
}
