package algorithms

// FindAllInverses computes all inverses mod m in O(n).
func FindAllInverses(m int64) []int64 {
	// Let r[i] be the inverse if i mod m
	// Then: r[i] = -(m / i) * r[m mod i]
	// Proof:
	// m mod i = m - (m / i) * i
	// m mod i = -(m / i) * i mod m | *r[i] *r[m mod i]
	// r[i] = -(m / i) * r[m mod i]
	// qed
	// This recurrent relationship actually works because m mod i < i.
	r := make([]int64, m)
	r[1] = 1
	for i := int64(2); i < m; i++ {
		r[i] = -(m / i) * r[m%i]
		if r[i] < 0 {
			r[i] += m
		}
	}
	return r
}
