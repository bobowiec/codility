package prefixsums

import "testing"

func TestGenomicRangeQuery(t *testing.T) {
	var table = []struct {
		S   string
		P   []int
		Q   []int
		ans []int
	}{
		{"CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}, []int{2, 4, 1}},
	}

	for _, test := range table {
		res := GenomicRangeQuery(test.S, test.P, test.Q)
		if !Equal(res, test.ans) {
			t.Errorf("GenomicRangeQuery was incorrect, got: %v, expected: %v\n", res, test.ans)
		}
	}
}

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
