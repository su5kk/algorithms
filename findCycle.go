package algorithms

// IsCyclic checks if a given graph is cyclic.
// Can be used as a helper for topological sort.
func IsCyclic(g [][]int64) bool {
	colors := make([]int64, len(g))
	for i := int64(0); i < int64(len(g)); i++ {
		if dfsColors(i, g, colors) {
			return true
		}
	}
	return false
}

func dfsColors(v int64, g [][]int64, colors []int64) bool {
	colors[v] = 1
	for i := int64(0); i < int64(len(g[v])); i++ {
		to := g[v][i]
		if colors[to] == 0 {
			if dfsColors(to, g, colors) {
				return true
			}
		} else if colors[to] == 1 {
			return true
		}
	}
	colors[v] = 2
	return false
}
