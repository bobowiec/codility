package fibonacci

import "testing"

func TestSumOfFibo(t *testing.T) {
	var table = []struct {
		numbers []int
		m       int
		ans     []bool
	}{
		{[]int{10, 15, 21, 25}, 1000000, []bool{true, true, true, false}},
	}

	for _, test := range table {
		result := SumOfFibo(test.numbers, test.m)
		for i, r := range result {
			if r != test.ans[i] {
				t.Errorf("SumOfFibo is incorrect, number: %d got: %t expected: %t\n", test.numbers[i], r, test.ans[i])
			}
		}
	}
}
