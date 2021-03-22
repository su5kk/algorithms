package algorithms

// Phi computes Euler totient function.
// we use O(sqrt(n)) factorization method.
func Phi(n int64) int64 {
	result := n
	// searching for divisors
	for i := int64(2); i*i <= n; i++ {
		// if we found a divisor.
		if n%i == 0 {
			// then we need to find all powers of this divisor.
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}
