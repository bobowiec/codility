package maxsliceproblem

import "testing"

func TestMaxProfit(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{23171, 21011, 21123, 21366, 21013, 21367}, 356},
	}

	for _, test := range table {
		result := MaxProfit(test.A)
		if result != test.ans {
			t.Errorf("MaxProfit is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
