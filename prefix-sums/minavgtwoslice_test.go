package prefixsums

import "testing"

func TestMinAvgTwoSlice(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{4, 2, 2, 5, 1, 5, 8}, 1},
		//{[]int {10, 10, -1, 2, 4, -1, 2, -1}, 5},
	}

	for _, test := range table {
		result := MinAvgTwoSlice(test.A)
		if result != test.ans {
			t.Errorf("MinAvgTwoSlice was incorrect, got: %d, expected: %d\n", result, test.ans)
		}
	}
}
