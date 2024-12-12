package datastructures

import "testing"

func TestArrListLen(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 4, -1, 3, 2}, 4},
	}

	for _, test := range table {
		result := ArrListLen(test.A)
		if result != test.ans {
			t.Errorf("ArrListLen is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
