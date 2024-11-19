package counting

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

func TestMaxCounters(t *testing.T) {
	var cases = []struct {
		N int
		A []int
		R []int
	}{
		{5, []int{3, 4, 4, 6, 1, 4, 4}, []int{3, 2, 2, 4, 2}},
		{4, []int{3, 4, 4, 5, 1, 4, 4, 5, 2, 3}, []int{4, 5, 5, 4}},
	}
	for _, test := range cases {
		result := MaxCounters(test.N, test.A)
		if !Equal(result, test.R) {
			t.Errorf("Max counters was incorrect, got: %v, expected: %v", result, test.R)
		}
	}
}
