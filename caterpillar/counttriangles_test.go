package caterpillar

import "testing"

func TestCountTriangles(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{10, 2, 5, 1, 8, 12}, 4},
		{[]int{4, 6, 3, 7}, 3},
		{[]int{10, 21, 22, 100, 101, 200, 300}, 6},
	}

	for _, test := range table {
		result := CountTriangles(test.A)
		if result != test.ans {
			t.Errorf("CountTriangle is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
