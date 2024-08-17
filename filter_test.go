package it_test

import (
	"fmt"

	"github.com/kyminbb/it"
)

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
