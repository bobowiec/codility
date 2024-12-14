package algorithmicskills

import "testing"

func TestUniqueNumber(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{4, 10, 5, 4, 2, 10}, 5},
		{[]int{1, 4, 3, 3, 1, 2}, 4},
		{[]int{6, 4, 4, 6}, -1},
	}

	for _, test := range table {
		result := FirstUnique(test.A)
		if result != test.ans {
			t.Errorf("UniqueNumber is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
