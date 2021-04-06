package algorithms

// Reverse apparently go does not have a reverse so....
// uses 2 pointers method
func Reverse(a []int64) []int64 {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
