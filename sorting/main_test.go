package main

import (
	"testing"
)

func TestSecondHighest(t *testing.T) {
	slice := []int{2, 33, 8, 1, 10, 12}
	result := SecondHighest(slice)
	if result != 12 {
		t.Errorf("expecting 8, but got %d", result)
	}
}
