package caterpillar

import "testing"

func TestMinAbsSumOfTwo(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 4, -3}, 1},
		{[]int{-8, 4, 5, -10, 3}, 3},
		{[]int{1000000000, 1000000000}, 2000000000},
	}

	for _, test := range table {
		result := MinAbsSumOfTwo(test.A)
		if result != test.ans {
			t.Errorf("MinAbsSumOfTwo is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
