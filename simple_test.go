package it_test

import (
	"testing"

	"github.com/kyminbb/it"
)

func TestAdd(t *testing.T) {
	if it.Add(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
}
