package counting

import "testing"

func TestPermCheck(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{4, 1, 3, 2}, 1},
		{[]int{4, 1, 3}, 0},
		{[]int{1, 4}, 0},
		{[]int{2, 2, 2}, 0},
		{[]int{9, 5, 7, 3, 2, 7, 3, 1, 10, 8}, 0},
		{[]int{1}, 1},
	}
	for _, test := range table {
		result := PermCheck(test.A)
		if result != test.ans {
			t.Errorf("PermCheck was incorrect, got: %d but expected: %d", result, test.ans)
		}
	}
}
