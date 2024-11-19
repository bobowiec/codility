package counting

import "testing"

func TestFrogRiverOne(t *testing.T) {
	var cases = []struct {
		A   []int
		X   int
		ans int
	}{
		{[]int{1, 3, 1, 4, 2, 3, 5, 4}, 5, 6},
		{[]int{2, 2, 2, 2, 2}, 2, -1},
		{[]int{1, 3, 1, 3, 2, 1, 3}, 3, 4},
	}
	for _, test := range cases {
		result := FrogRiverOne(test.X, test.A)
		if result != test.ans {
			t.Errorf("FrogRiverOne was incorrect, got: %d but expected: %d", result, test.ans)
		}
	}
}
