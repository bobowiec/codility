package maxsliceproblem

import "testing"

func TestMaxDoubleSliceSum(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{3, 2, 6, -1, 4, 5, -1, 2}, 17},
		{[]int{0, 10, -5, -2, 0}, 10},
	}

	for _, test := range table {
		result := MaxDoubleSliceSum(test.A)
		if result != test.ans {
			t.Errorf("MaxDoubleSliceSum is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
