package datastructures

/*
https://app.codility.com/programmers/trainings/7/count_bounded_slices/

An integer K and a non-empty array A consisting of N integers are given.

A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A.

A bounded slice is a slice in which the difference between the maximum and minimum values in the slice is less than or equal to K. More precisely it is a slice, such that max(A[P], A[P + 1], ..., A[Q]) − min(A[P], A[P + 1], ..., A[Q]) ≤ K.

The goal is to calculate the number of bounded slices.

For example, consider K = 2 and array A such that:

    A[0] = 3
    A[1] = 5
    A[2] = 7
    A[3] = 6
    A[4] = 3
There are exactly nine bounded slices: (0, 0), (0, 1), (1, 1), (1, 2), (1, 3), (2, 2), (2, 3), (3, 3), (4, 4).

Write a function:

func Solution(K int, A []int) int

that, given an integer K and a non-empty array A of N integers, returns the number of bounded slices of array A.

If the number of bounded slices is greater than 1,000,000,000, the function should return 1,000,000,000.

For example, given:

    A[0] = 3
    A[1] = 5
    A[2] = 7
    A[3] = 6
    A[4] = 3
the function should return 9, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
K is an integer within the range [0..1,000,000,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].
*/

func CountBoundedSlices(K int, A []int) int {
	const maxINT = 1000000000
	N := len(A)

	maxQ := make([]int, N+1)
	posmaxQ := make([]int, N+1)
	minQ := make([]int, N+1)
	posminQ := make([]int, N+1)

	firstMax, lastMax := 0, -1
	firstMin, lastMin := 0, -1
	j, result := 0, 0

	for i := 0; i < N; i++ {
		for j < N {
			// added new maximum element
			for lastMax >= firstMax && maxQ[lastMax] <= A[j] {
				lastMax--
			}
			lastMax++
			maxQ[lastMax] = A[j]
			posmaxQ[lastMax] = j
			// added new minimum element
			for lastMin >= firstMin && minQ[lastMin] >= A[j] {
				lastMin--
			}
			lastMin++
			minQ[lastMin] = A[j]
			posminQ[lastMin] = j

			if maxQ[firstMax]-minQ[firstMin] <= K {
				j += 1
			} else {
				break
			}
		}
		result += (j - i)
		if result >= maxINT {
			return maxINT
		}
		if posminQ[firstMin] == i {
			firstMin++
		}
		if posmaxQ[firstMax] == i {
			firstMax++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CountBoundedSlicesBruteForce(K int, A []int) int {
	N := len(A)
	result := 0
	for i := range A {
		minimum := A[i]
		maximum := A[i]
		for j := i; j < N; j++ {
			maximum = max(maximum, A[j])
			minimum = min(minimum, A[j])
			if maximum-minimum <= K {
				result += 1
				if result == 1000000000 {
					return result
				}
			} else {
				break
			}
		}
	}
	return result
}
