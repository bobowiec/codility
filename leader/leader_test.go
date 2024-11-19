package leader

import "testing"

func TestLeader(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{6, 8, 4, 6, 8, 6, 6}, 6},
	}
	for _, test := range table {
		result := slowLeader(test.A)
		if result != test.ans {
			t.Errorf("slowLeader is incorrect, got: %d expected: %d\n", result, test.ans)
		}

		result = fastLeader(test.A)
		if result != test.ans {
			t.Errorf("fastLeader is incorrect, got: %d expected: %d\n", result, test.ans)
		}

		result = goldenLeader(test.A)
		if result != test.ans {
			t.Errorf("goldenLeader is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
