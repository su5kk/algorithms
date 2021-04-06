package algorithms

// TopologicalSort sorts graph via dfs.
func TopologicalSort(g [][]int64) []int64 {
	ans := make([]int64, 0)
	visited := make([]bool, len(g))
	for i := int64(0); i < int64(len(g)); i++ {
		if !visited[i] {
			DFS(i, visited, &ans, g)
		}
	}
	return Reverse(ans)
}
