package binarysearch

/*
https://app.codility.com/programmers/lessons/14-binary_search_algorithm/nailing_planks/

You are given two non-empty arrays A and B consisting of N integers. These arrays represent N planks. More precisely, A[K] is the start and B[K] the end of the K−th plank.

Next, you are given a non-empty array C consisting of M integers. This array represents M nails. More precisely, C[I] is the position where you can hammer in the I−th nail.

We say that a plank (A[K], B[K]) is nailed if there exists a nail C[I] such that A[K] ≤ C[I] ≤ B[K].

The goal is to find the minimum number of nails that must be used until all the planks are nailed. In other words, you should find a value J such that all planks will be nailed after using only the first J nails. More precisely, for every plank (A[K], B[K]) such that 0 ≤ K < N, there should exist a nail C[I] such that I < J and A[K] ≤ C[I] ≤ B[K].

For example, given arrays A, B such that:

    A[0] = 1    B[0] = 4
    A[1] = 4    B[1] = 5
    A[2] = 5    B[2] = 9
    A[3] = 8    B[3] = 10
four planks are represented: [1, 4], [4, 5], [5, 9] and [8, 10].

Given array C such that:

    C[0] = 4
    C[1] = 6
    C[2] = 7
    C[3] = 10
    C[4] = 2
if we use the following nails:

0, then planks [1, 4] and [4, 5] will both be nailed.
0, 1, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
0, 1, 2, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
0, 1, 2, 3, then all the planks will be nailed.
Thus, four is the minimum number of nails that, used sequentially, allow all the planks to be nailed.

Write a function:

func Solution(A []int, B []int, C []int) int

that, given two non-empty arrays A and B consisting of N integers and a non-empty array C consisting of M integers, returns the minimum number of nails that, used sequentially, allow all the planks to be nailed.

If it is not possible to nail all the planks, the function should return −1.

For example, given arrays A, B, C such that:

    A[0] = 1    B[0] = 4
    A[1] = 4    B[1] = 5
    A[2] = 5    B[2] = 9
    A[3] = 8    B[3] = 10

    C[0] = 4
    C[1] = 6
    C[2] = 7
    C[3] = 10
    C[4] = 2
the function should return 4, as explained above.

Write an efficient algorithm for the following assumptions:

N and M are integers within the range [1..30,000];
each element of arrays A, B and C is an integer within the range [1..2*M];
A[K] ≤ B[K].
*/

func canNailAllPlanksWithPrefixSum(A, B, C []int, k int) bool {
	maxPosition := 0
	for _, pos := range C[:k] {
		if pos > maxPosition {
			maxPosition = pos
		}
	}

	// Initialize nail positions and prefix sum arrays
	prefixSums := make([]int, maxPosition+2)
	for i := 0; i < k; i++ {
		prefixSums[C[i]]++
	}

	// Build prefix sum array
	for i := 1; i < len(prefixSums); i++ {
		prefixSums[i] += prefixSums[i-1]
	}

	// Check each plank if it can be nailed
	for i := 0; i < len(A); i++ {
		if prefixSums[B[i]] == prefixSums[A[i]-1] {
			return false
		}
	}
	return true
}

func NailingPlanks(A []int, B []int, C []int) int {
	m := len(C)
	left, right := 1, m
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if canNailAllPlanksWithPrefixSum(A, B, C, mid) {
			right = mid - 1
			result = mid
		} else {
			left = mid + 1
		}
	}
	return result
}

func NailingPlanksBruteForce(A []int, B []int, C []int) int {
	n := len(A)
	nailed := make([]bool, n)
	for i, v := range C {
		for k := range A {
			if nailed[k] {
				continue
			}
			if v >= A[k] && v <= B[k] {
				nailed[k] = true
			}
		}
		if allNailed(nailed) {
			return i + 1
		}
	}
	return -1
}

func allNailed(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}
