package algorithms

// LinearSearch returns id and the result of execution.
// id = last possible index and result = true iff element was found.
// id = -1 and result = false iff element was not found.
func LinearSearch(a []int, b int) (id int, result bool) {
	id = -1
	result = false
	for idx, elem := range a {
		if elem == b {
			id = idx
			result = true
		}
	}
	return id, result
}
