package algorithms

import (
	"testing"
)

func TestFibonacciSmallNum(t *testing.T) {
	n := int64(9)
	m := int64(1e9 + 7)
	want := int64(21)
	if Fib(n, m) != 21 {
		t.Fatalf(`Fib(n, m) = %v, want: %v`, Fib(n, m), want)
	}
}

func TestFibonacciMediumNum(t *testing.T) {
	n := int64(33)
	m := int64(1e9 + 7)
	want := int64(2178309)
	if Fib(n, m) != want {
		t.Fatalf(`Fib(n, m) = %v, want: %v`, Fib(n, m), want)
	}
}
