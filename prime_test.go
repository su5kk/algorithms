package algorithms

import "testing"

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		name     string
		input    int64
		expected bool
	}{
		{"smallest prime", 2, true},
		{"random prime", 3, true},
		{"neither prime nor composite", 1, false},
		{"random non-prime", 10, false},
		{"another random prime", 23, true},
		{"huge prime", 2147483647, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := IsPrime(test.input); output != test.expected {
				t.Errorf("For input: %d, expected: %v, but got: %v", test.input, test.expected, output)
			}
		})
	}

}
