package main

import (
	"testing"
)

func TestDups(t *testing.T) {
	slice := []int{2, 3, 2}
	result := Dups(slice)
	if result != true {
		t.Errorf("expecting true, got %t ", result)
	}
}
