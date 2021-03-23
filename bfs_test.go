package algorithms

import "testing"

func slicesEqual(a, b []int64) bool {
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
func TestFindShortestPathValid(t *testing.T) {
	graph := [][]int64{
		0: {1, 4},
		1: {0, 2},
		2: {1, 3, 4, 5},
		3: {2},
		4: {0, 2},
		5: {2},
	}
	to := int64(5)
	want := []int64{0, 1, 2, 5}
	wantAlt := []int64{0, 4, 2, 5}
	result := FindShortestPath(graph, 6, to)
	if !slicesEqual(result, want) {
		t.Fatalf(`FindShortestPath(graph, numberOfVertices, to) = %v, want: %v`, result, want)
		return
	}
	if !slicesEqual(result, wantAlt) && !slicesEqual(result, want) {
		t.Fatalf(`FindShortestPath(graph, numberOfVertices, to) = %v, wantAlt: %v`, result, wantAlt)
		return
	}
}

func TestFindShortestPathNoPath(t *testing.T) {
	graph := [][]int64{
		0: {1, 4},
		1: {0, 2},
		2: {1, 3, 4, 5},
		3: {2},
		4: {0, 2},
		5: {2},
		6: {6},
	}
	to := int64(6)
	want := []int64{}
	result := FindShortestPath(graph, 7, to)
	if !slicesEqual(result, want) {
		t.Fatalf(`FindShortestPath(graph, numberOfVertices, to) = %v, want: %v`, result, want)
		return
	}
}
