package algorithms

import "testing"

func TestSearchPathYES(t *testing.T) {
	graph := [][]int64{
		0: {1},
		1: {1},
		2: {2},
	}
	s := int64(0)
	v := int64(1)
	want := true
	if !SearchPath(graph, s, v) {
		t.Fatalf(`SearchPath(graph, s, t) = %v, want: %v`, SearchPath(graph, s, v), want)
	}
}

func TestSearchPathNO(t *testing.T) {
	graph := [][]int64{
		0: {1},
		1: {1},
		2: {2},
	}
	s := int64(0)
	v := int64(2)
	want := false
	if SearchPath(graph, s, v) {
		t.Fatalf(`SearchPath(graph, s, t) = %v, want: %v`, SearchPath(graph, s, v), want)
	}
}
