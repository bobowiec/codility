package euclidean

import "testing"

func TestCommonPrimeDivisors(t *testing.T) {
	var table = []struct {
		A   []int
		B   []int
		ans int
	}{
		{[]int{15, 10, 3}, []int{75, 30, 5}, 1},
		{[]int{1}, []int{1}, 1},
	}

	for _, test := range table {
		result := CommonPrimeDivisors(test.A, test.B)
		if result != test.ans {
			t.Errorf("CommonPrimeDivisors is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
