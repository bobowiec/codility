package algorithmicskills

import "testing"

func TestArrayInversionCount(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{-1, 6, 3, 4, 7, 4}, 4},
		{[]int{10, 2, 1, 4, 1, -1, 19, 1}, 16},
	}

	for _, test := range table {
		result := ArrayInversionCount(test.A)
		if result != test.ans {
			t.Errorf("ArrayInversionCount is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
