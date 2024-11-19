package binarysearch

import "testing"

func TestBoards(t *testing.T) {
	var table = []struct {
		A   []int
		k   int
		ans int
	}{
		{[]int{0, 1, 0, 1, 1, 1, 0, 1}, 1, 7},
		{[]int{0, 1, 0, 1, 1, 1, 0, 1}, 2, 4},
		{[]int{0, 1, 0, 1, 1, 1, 0, 1}, 3, 3},
		{[]int{0, 1, 0, 1, 1, 1, 0, 1}, 4, 2},
		{[]int{0, 1, 0, 1, 1, 1, 0, 1}, 5, 1},
		{[]int{0, 1, 0, 1, 1, 1, 0, 1}, 6, -1},
	}

	for _, test := range table {
		result := Boards(test.A, test.k)
		if result != test.ans {
			t.Errorf("Boards is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
