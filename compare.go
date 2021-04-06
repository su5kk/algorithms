package algorithms

// Compare compares two int64 arrays.
func Compare(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for id := range a {
		if a[id] != b[id] {
			return false
		}
	}
	return true
}
