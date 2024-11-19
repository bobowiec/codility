package binarysearch

import "testing"

func TestMinMaxDivision(t *testing.T) {
	var table = []struct {
		K   int
		M   int
		A   []int
		ans int
	}{
		{3, 5, []int{2, 1, 5, 1, 2, 2, 2}, 6},
	}

	for _, test := range table {
		result := MinMaxDivision(test.K, test.M, test.A)
		if result != test.ans {
			t.Errorf("MinMaxDivision is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
