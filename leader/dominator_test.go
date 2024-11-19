package leader

import "testing"

func TestDominator(t *testing.T) {
	var table = []struct {
		A   []int
		ans []int
	}{
		{[]int{3, 4, 3, 2, 3, -1, 3, 3}, []int{0, 2, 4, 6, 7}},
		{[]int{6, 8, 4, 6, 8, 6, 6}, []int{0, 3, 5, 6}},
	}
	for _, test := range table {
		result := Dominator(test.A)
		found := false
		for _, v := range test.ans {
			if result == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Dominator is incorrect, got: %d expected: %v\n", result, test.ans)
		}

	}
}
