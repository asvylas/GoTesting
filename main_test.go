package main

import "testing"

func TestAddTwo(t *testing.T) {
	if addTwo(5, 5) != 10 {
		t.Error("Test failed, expected: 10")
	}
}
