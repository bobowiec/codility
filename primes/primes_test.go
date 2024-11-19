package primes

import "testing"

func TestCountDivisors(t *testing.T) {
	var table = []struct {
		n   int
		ans int
	}{
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 3},
	}

	for _, test := range table {
		result := countDivisors(test.n)
		if result != test.ans {
			t.Errorf("CountDivisors is incorrect, n=%d got: %d expected: %d\n", test.n, result, test.ans)
		}
	}
}

func TestPrimality(t *testing.T) {
	var table = []struct {
		n   int
		ans bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
	}

	for _, test := range table {
		result := primality(test.n)
		if result != test.ans {
			t.Errorf("Primality is incorrect, n=%d got: %v expected: %v\n", test.n, result, test.ans)
		}
	}
}

func TestCoins(t *testing.T) {
	var table = []struct {
		n   int
		ans int
	}{
		{9, 3}, //?
	}

	for _, test := range table {
		result := coins(test.n)
		if result != test.ans {
			t.Errorf("Coins is incorrect, n=%d got: %v expected: %v\n", test.n, result, test.ans)
		}
	}
}
