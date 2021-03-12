package algorithms

// helper function
func merge(a, b []int) []int {
	result := []int{}
	i, j := 0, 0
	// Two pointers method.
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	// In case arrays have different size, we add the remaining items.
	result = append(result, b[j:]...)
	result = append(result, a[i:]...)
	return result
}

// SortMerges uses Divide and Conquer algorithm technique to sort arrays.
func SortMerge(a []int) []int {
	if len(a) > 1 {
		rightArray := a[len(a)/2:]
		leftArray := a[:len(a)/2]
		return merge(SortMerge(leftArray), SortMerge(rightArray))
	}
	return a
}
