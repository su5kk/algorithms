package algorithms

// DigitalRoot returns a digital root of a number with O(1) complexity.
func DigitalRoot(a int) int {
	return 1 + ((a - 1) % 9)
}
