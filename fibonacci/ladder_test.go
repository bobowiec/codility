package fibonacci

import "testing"

func TestLadder(t *testing.T) {
	var table = []struct {
		A   []int
		B   []int
		ans []int
	}{
		{[]int{4, 4, 5, 5, 1}, []int{3, 2, 4, 3, 1}, []int{5, 1, 8, 0, 1}},
	}

	for _, test := range table {
		result := Ladder(test.A, test.B)
		for i, r := range result {
			if r != test.ans[i] {
				t.Errorf("Ladder is incorrect, got: %v expected: %v\n", result, test.ans)
			}
		}
	}
}
