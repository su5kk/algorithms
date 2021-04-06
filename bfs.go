package algorithms

import (
	"container/list"
)

// FindShortestPath finds the shortest path :O
func FindShortestPath(graph [][]int64, numberOfVertices, to int64) []int64 {
	path := []int64{}
	used, parents := bfs(graph, numberOfVertices, 0)
	if !used[to] {
		return path
	}
	for i := to; i != -1; i = parents[i] {
		path = append(path, i)
	}

	return Reverse(path)
}

func bfs(graph [][]int64, numberOfVertices, start int64) ([]bool, []int64) {
	queue := list.New()
	queue.PushBack(start)
	// is needed to keep in mind what vertices we have already travelled to.
	used := make([]bool, numberOfVertices)
	// optional: needed to restore the info about the shortest paths.
	pathsLength := make([]int64, numberOfVertices)
	// optional: needed to restore the info about the shortest paths.
	parents := make([]int64, numberOfVertices)
	used[start] = true
	parents[start] = -1 // default value.
	for queue.Len() != 0 {
		current := queue.Front()
		v := current.Value.(int64)
		queue.Remove(queue.Front())
		for i := 0; i < len(graph[v]); i++ {
			to := graph[v][i]
			if !used[to] {
				used[to] = true
				queue.PushBack(to)
				pathsLength[to] = pathsLength[v] + 1
				parents[to] = v
			}
		}
	}
	return used, parents
}
