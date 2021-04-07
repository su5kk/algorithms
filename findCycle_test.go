package algorithms

import "testing"

func TestIsCyclic(t *testing.T) {
	testCases := []struct {
		desc     string
		input    [][]int64
		expected bool
	}{
		{
			desc: "Cyclic graph",
			input: [][]int64{
				0: {1},
				1: {2},
				2: {0},
			},
			expected: true,
		},
		{
			desc: "Acyclic graph",
			input: [][]int64{
				0: {1, 2},
				1: {3, 4},
				2: {},
				3: {},
				4: {},
			},
			expected: false,
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if output := IsCyclic(test.input); output != test.expected {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.input, test.expected, output)
			}
		})
	}
}
