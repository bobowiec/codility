package counting

/**
This is a demo task.

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/

// MissingInteger finds the smallest positive integer that does not occur in a given sequence.
func MissingInteger(A []int) int {
	result := 1
	seen := make(map[int]bool)
	for _, v := range A {
		if v > 0 && !seen[v] {
			seen[v] = true
		}
	}
	for i := 1; i <= len(seen); i++ {
		if !seen[result] {
			return result
		}
		result++
	}
	return result
}
