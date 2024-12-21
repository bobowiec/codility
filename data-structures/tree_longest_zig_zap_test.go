package datastructures

import "testing"

func TestLongestZigZag(t *testing.T) {
	root := &Tree{X: 5}
	root.L = &Tree{X: 3}
	root.L.L = &Tree{X: 20}
	root.L.L.L = &Tree{X: 6}
	root.R = &Tree{X: 10}
	root.R.L = &Tree{X: 1}
	root.R.R = &Tree{X: 15}
	root.R.R.L = &Tree{X: 30}
	root.R.R.L.R = &Tree{X: 9}
	root.R.R.R = &Tree{X: 8}

	var table = []struct {
		node *Tree
		ans  int
	}{
		{root, 2},
	}

	for _, test := range table {
		result := TreeLongestZigZag(test.node)
		if result != test.ans {
			t.Errorf("TreeLongestZigZag is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
