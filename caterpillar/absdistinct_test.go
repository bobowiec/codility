package caterpillar

import "testing"

func TestAbsDistinct(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{-5, -3, -1, 0, 3, 6}, 5},
	}

	for _, test := range table {
		result := AbsDistinct(test.A)
		if result != test.ans {
			t.Errorf("AbsDistinct is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
