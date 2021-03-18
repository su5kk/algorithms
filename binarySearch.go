package algorithms

// BinarySearch returns index of an element, if it is found.
func BinarySearch(a []int64, x int64) int64 {
	var l, r, m int64
	l = 0
	r = int64(len(a))
	m = -1
	for l < r {
		m := l + ((r - l) >> 1)
		if a[m] == x {
			return m
		} else if a[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return m
}
