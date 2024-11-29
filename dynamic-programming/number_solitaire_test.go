package dynamicprogramming

import "testing"

func TestNumberSolitaire(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, -2, 0, 9, -1, -2}, 8},
		{[]int{1, -2, 4, 3, -1, -3, -7, 4, -9}, 3},
	}

	for _, test := range table {
		result := NumberSolitaireOpt(test.A)
		if result != test.ans {
			t.Errorf("NumberSolitaire is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
