package dynamicprogramming

import "testing"

func TestMinAbsSum(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 5, 2, -2}, 0},
		{[]int{7}, 7},
		{[]int{7, -6}, 1},
		{[]int{3, 1}, 2},
	}

	for _, test := range table {
		result := MinAbsSum(test.A)
		if result != test.ans {
			t.Errorf("MinAbsSum is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
