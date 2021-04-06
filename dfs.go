package algorithms

// SearchPath does what you think it does.
func SearchPath(graph [][]int64, s, t int64) bool {
	visited := make([]bool, len(graph))
	return path(s, t, visited, graph)
}

func path(u, t int64, visited []bool, graph [][]int64) bool {
	if u == t {
		return true
	}
	visited[u] = true
	for _, neighbour := range graph[u] {
		if !visited[neighbour] {
			if path(neighbour, t, visited, graph) {
				return true
			}
		}
	}
	return false
}

// DFS is dfs.
func DFS(v int64, visited []bool, ans *[]int64, g [][]int64) {
	// invariant.
	visited[v] = true
	for i := 0; i < len(g[v]); i++ {
		to := g[v][i]
		if !visited[to] {
			DFS(to, visited, ans, g)
		}
	}
	*ans = append(*ans, v)
}
