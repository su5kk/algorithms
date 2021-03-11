package algorithms

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// GCD returns gcd of a and b.
func GCD(a, b int) int {
	a1 := max(a, b)
	b1 := min(a, b)
	for b1 != 0 {
		a1, b1 = b1, a1%b1
	}
	return a1
}
