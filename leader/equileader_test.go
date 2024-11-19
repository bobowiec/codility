package leader

import "testing"

func TestEquiLeader(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{4, 3, 4, 4, 4, 2}, 2},
		{[]int{0, 0}, 1},
	}

	for _, test := range table {
		result := EquiLeader(test.A)
		if result != test.ans {
			t.Errorf("EquiLeader is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
