package sorting

import "testing"

func TestDistinct(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{2, 1, 1, 2, 3, 1}, 3},
	}

	for _, test := range table {
		result := Distinct(test.A)
		result2 := DistinctBySort(test.A)

		if result != test.ans {
			t.Errorf("Distinct was incorrect, got: %d, expected %d\n", result, test.ans)
		}

		if result2 != test.ans {
			t.Errorf("DistinctBySort was incorrect, got: %d, expected %d\n", result2, test.ans)
		}
	}
}
