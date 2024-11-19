package euclidean

import "testing"

func TestChocolatesByNumbers(t *testing.T) {
	var table = []struct {
		N   int
		M   int
		ans int
	}{
		{10, 4, 5},
		{947853, 4453, 947853},
	}

	for _, test := range table {
		result := ChocolatesByNumbers(test.N, test.M)
		if result != test.ans {
			t.Errorf("ChocolatesByNumbers is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
