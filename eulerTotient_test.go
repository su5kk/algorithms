package algorithms

import "testing"

func TestPhiPrime(t *testing.T) {
	n := int64(7)
	want := 6
	if Phi(n) != 6 {
		t.Fatalf(`Phi(n) = %v, want: %v`, Phi(n), want)
	}
}

func TestPhiComposite(t *testing.T) {
	n := int64(8)
	want := 4
	if Phi(n) != 4 {
		t.Fatalf(`Phi(n) = %v, want: %v`, Phi(n), want)
	}
}
