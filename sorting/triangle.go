/*
https://app.codility.com/programmers/lessons/6-sorting/triangle/

An array A consisting of N integers is given. A triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:

A[P] + A[Q] > A[R],
A[Q] + A[R] > A[P],
A[R] + A[P] > A[Q].
For example, consider array A such that:

	A[0] = 10    A[1] = 2    A[2] = 5
	A[3] = 1     A[4] = 8    A[5] = 20

Triplet (0, 2, 4) is triangular.

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns 1 if there exists a triangular triplet for this array and returns 0 otherwise.

For example, given array A such that:

	A[0] = 10    A[1] = 2    A[2] = 5
	A[3] = 1     A[4] = 8    A[5] = 20

the function should return 1, as explained above. Given array A such that:

	A[0] = 10    A[1] = 50    A[2] = 5
	A[3] = 1

the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
*/
package sorting

import "sort"

func Triangle(A []int) int {
	sort.Ints(A)
	N := len(A)
	for i := 0; i <= N-3; i++ {
		if A[i]+A[i+1] > A[i+2] && A[i]+A[i+2] > A[i+1] && A[i+1]+A[i+2] > A[i] {
			return 1
		}
	}
	return 0
}

func Triangle2(A []int) int {
	N := len(A)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if A[i]+A[j] > A[k] &&
					A[i]+A[k] > A[j] &&
					A[j]+A[k] > A[i] {
					return 1
				}
			}
		}
	}
	return 0
}
