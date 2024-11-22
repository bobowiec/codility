package greedy

import "testing"

func TestGreedyCoinChanging(t *testing.T) {
	var table = []struct {
		M   []int
		k   int
		ans []int
	}{
		{[]int{1, 2, 5}, 6, []int{5, 1, 2, 0, 1, 1}},
	}

	for _, test := range table {
		result := GreedyCoinChanging(test.M, test.k)
		if len(result) != len(test.ans) {
			t.Errorf("GreedyCoinChanging is incorrect, got: %v expected: %v\n", result, test.ans)
		}
		for i := range result {
			if result[i] != test.ans[i] {
				t.Errorf("GreedyCoinChanging is incorrect, got: %v expected: %v\n", result, test.ans)
			}
		}
	}
}
