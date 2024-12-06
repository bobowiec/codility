package codingskills

import "testing"

func TestParkingBill(t *testing.T) {
	var table = []struct {
		E   string
		L   string
		ans int
	}{
		{"10:00", "13:21", 17},
		{"09:42", "11:42", 9},
	}

	for _, test := range table {
		result := ParkingBill(test.E, test.L)
		if result != test.ans {
			t.Errorf("ParkingBill is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
