package sieve

import "testing"

func TestCountNonDivisible(t *testing.T) {
	var table = []struct {
		A   []int
		ans []int
	}{
		{[]int{3, 1, 2, 3, 6}, []int{2, 4, 3, 2, 0}},
		{[]int{2, 4}, []int{1, 0}},
		{[]int{4, 4}, []int{0, 0}},
	}

	for _, test := range table {
		results := CountNonDivisible(test.A)
		for i, k := range results {
			if k != test.ans[i] {
				t.Errorf("CoundNonDivisible is incorrect, got: %v expected:%v\n", results, test.ans)
			}
		}
	}
}
