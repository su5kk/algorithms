package algorithms

// IsPrime does what you think it does.
func IsPrime(n int64) bool {
	if n == 1 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
