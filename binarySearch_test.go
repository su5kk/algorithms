package algorithms

import (
	"testing"
)

func TestBinarySearchElementInArray(t *testing.T) {
	a := []int64{1, 2, 3, 4}
	x := int64(3)
	want := int64(2)
	if BinarySearch(a, x) != 2 {
		t.Fatalf(`BinarySearch(a, x) = %v, want: %v`, BinarySearch(a, x), want)
	}
}

func TestBinarySearchElementNotInArray(t *testing.T) {
	a := []int64{1, 2, 3, 4}
	x := int64(5)
	want := int64(-1)
	if BinarySearch(a, x) != -1 {
		t.Fatalf(`BinarySearch(a, x) = %v, want: %v`, BinarySearch(a, x), want)
	}
}
