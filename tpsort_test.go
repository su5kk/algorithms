package algorithms

import "testing"

func TestTopologicalSort(t *testing.T) {
	testestases := []struct {
		desc     string
		input    [][]int64
		expected []int64
	}{
		{
			desc: "basic graph",
			input: [][]int64{
				0: {1, 2},
				1: {1},
				2: {1},
			},
			expected: []int64{0, 2, 1},
		},
	}
	for _, test := range testestases {
		t.Run(test.desc, func(t *testing.T) {
			if output := TopologicalSort(test.input); !Compare(output, test.expected) {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.input, test.expected, output)
			}
		})
	}
}
