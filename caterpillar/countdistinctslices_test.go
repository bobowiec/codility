package caterpillar

import "testing"

func TestCountDistinctSlices(t *testing.T) {
	var table = []struct {
		M   int
		A   []int
		ans int
	}{
		{6, []int{3, 4, 5, 5, 2}, 9},
	}

	for _, test := range table {
		result := CountDistinctSlices(test.M, test.A)
		if result != test.ans {
			t.Errorf("CountDistinctSlices is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
