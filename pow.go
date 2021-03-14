package algorithms

// Pow does what you think it does.
func Pow(a, x int64) int64 {
	var result int64
	result = 1
	for x > 0 {
		// for any x, x can be represented as the sum of 2 to some power.
		// for example, 13 = 2^3 + 2^2 + 2^0 => 13 = 1101 base 2
		// so, a^13 = a^8 + a^4 + a^1
		// as we can see, we only need to multiply the result iff current bit is 1
		// and skip otherwise.
		// a is also squares each time regardless of the current bit.
		if x&1 == 1 {
			result *= a
		}
		a *= a
		x >>= 1
	}
	return result
}
