package it_test

import (
	"fmt"

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
