package stacksqueues

import "testing"

func TestStoneWall(t *testing.T) {
	var table = []struct {
		H   []int
		ans int
	}{
		{[]int{8, 8, 5, 7, 9, 8, 7, 4, 8}, 7},
		{[]int{2, 3, 2, 1}, 3},
	}

	for _, test := range table {
		result := StoneWall(test.H)
		if result != test.ans {
			t.Errorf("Stonewall is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
