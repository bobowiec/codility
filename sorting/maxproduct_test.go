package sorting

import "testing"

func TestMaxProductOfThree(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{-3, 1, 2, -2, 5, 6}, 60},
		{[]int{-5, 5, -5, 4}, 125},
		{[]int{-6, -4, 3, 4, 5}, 120},
	}

	for _, test := range table {
		result := MaxProductOfThree(test.A)
		if result != test.ans {
			t.Errorf("MaxProductOfThree is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
