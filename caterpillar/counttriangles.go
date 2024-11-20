package caterpillar

import (
	"sort"
)

/*
https://app.codility.com/programmers/lessons/15-caterpillar_method/count_triangles/

An array A consisting of N integers is given. A triplet (P, Q, R) is triangular if it is possible to build a triangle with sides of lengths A[P], A[Q] and A[R]. In other words, triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:

A[P] + A[Q] > A[R],
A[Q] + A[R] > A[P],
A[R] + A[P] > A[Q].
For example, consider array A such that:

  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 12
There are four triangular triplets that can be constructed from elements of this array, namely (0, 2, 4), (0, 2, 5), (0, 4, 5), and (2, 4, 5).

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the number of triangular triplets in this array.

For example, given array A such that:

  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 12
the function should return 4, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..1,000];
each element of array A is an integer within the range [1..1,000,000,000].
*/

func CountTriangles(A []int) int {
	n := len(A)
	sort.Ints(A)
	count := 0
	for i := n - 1; i >= 0; i-- {
		back, front := 0, i-1
		for back < front {
			if A[back]+A[front] > A[i] {
				count += front - back
				front--
			} else {
				back++
			}
		}
	}
	return count
}

func CountTrianglesBruteForce(A []int) int {
	count := 0
	n := len(A)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if A[i]+A[j] > A[k] && A[i]+A[k] > A[j] && A[j]+A[k] > A[i] {
					count++
				}
			}
		}
	}
	return count
}
