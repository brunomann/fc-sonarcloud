package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(2, 2)

	if result != 4 {
		t.Error("The result must be 4")
	}
}
