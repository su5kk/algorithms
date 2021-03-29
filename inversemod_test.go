package algorithms

import "testing"

func TestInverse(t *testing.T) {
	var tests = []struct {
		name     string
		input    int64
		expected []int64
	}{
		{"Small set", 7, []int64{0, 1, 4, 5, 2, 3, 6}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := FindAllInverses(test.input); !compare(output, test.expected) {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.input, test.expected, output)
			}
		})
	}
}

func compare(a, b []int64) bool {
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
