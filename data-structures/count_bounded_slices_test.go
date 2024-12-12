package datastructures

import "testing"

func TestCountBoundedSlices(t *testing.T) {
	var table = []struct {
		K   int
		A   []int
		ans int
	}{
		{2, []int{3, 5, 7, 6, 3}, 9},
	}

	for _, test := range table {
		result := CountBoundedSlices(test.K, test.A)
		if result != test.ans {
			t.Errorf("CountBoundedSlices is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
