package maxsliceproblem

import "testing"

func TestMaxSliceSum(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{3, 2, -6, 4, 0}, 5},
		{[]int{-10}, -10},
		{[]int{-2, 1}, 1},
	}

	for _, test := range table {
		result := MaxSliceSum(test.A)
		if result != test.ans {
			t.Errorf("MaxSliceSum is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
