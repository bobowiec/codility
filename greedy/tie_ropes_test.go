package greedy

import "testing"

func TestTieRopes(t *testing.T) {
	var table = []struct {
		K   int
		A   []int
		ans int
	}{
		{4, []int{1, 2, 3, 4, 1, 1, 3}, 3},
		{2, []int{1}, 0},
	}

	for _, test := range table {
		result := TieRopes(test.K, test.A)
		if result != test.ans {
			t.Errorf("TieRopes is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
