package main

import (
	"testing"
)

func TestDups(t *testing.T) {
	slice := []int{2, 3, 2}
	result := Dups(slice)
	if result != true {
		t.Errorf("expecting true, got %t", result)
	}

	slice = []int{3, 4, 5, 6}
	result = Dups(slice)
	if result != false {
		t.Errorf("expecting false, got %t", result)
	}

	slice = []int{2, 3, 7, 4, 11, 100, 99, 12, 34, 899}
	result = Dups(slice)
	if result != false {
		t.Errorf("expecting false, got %t", result)
	}

}
