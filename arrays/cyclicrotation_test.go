package arrays

import "testing"

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCyclicRotate(t *testing.T) {
	var tests = []struct {
		A []int
		K int
		R []int
	}{
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		output := CyclicRotate(test.A, test.K)
		if !Equal(output, test.R) {
			t.Errorf("Cyclic rotation was incorrect, got: %v, expected: %v", output, test.R)
		}
	}
}
