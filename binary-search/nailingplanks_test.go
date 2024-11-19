package binarysearch

import "testing"

func TestNailingPlanks(t *testing.T) {
	var table = []struct {
		A   []int
		B   []int
		C   []int
		ans int
	}{
		{[]int{1, 4, 5, 8}, []int{4, 5, 9, 10}, []int{4, 6, 7, 10, 2}, 4},
		{[]int{2}, []int{2}, []int{1}, -1},
	}

	for _, test := range table {
		result := NailingPlanks(test.A, test.B, test.C)
		if result != test.ans {
			t.Errorf("NailingPlanks is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
