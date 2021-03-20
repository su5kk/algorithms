package algorithms

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
