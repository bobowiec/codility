package binarysearch

/*
https://app.codility.com/programmers/lessons/14-binary_search_algorithm/min_max_division/

You are given integers K, M and a non-empty array A consisting of N integers. Every element of the array is not greater than M.

You should divide this array into K blocks of consecutive elements. The size of the block is any integer between 0 and N. Every element of the array should belong to some block.

The sum of the block from X to Y equals A[X] + A[X + 1] + ... + A[Y]. The sum of empty block equals 0.

The large sum is the maximal sum of any block.

For example, you are given integers K = 3, M = 5 and array A such that:

  A[0] = 2
  A[1] = 1
  A[2] = 5
  A[3] = 1
  A[4] = 2
  A[5] = 2
  A[6] = 2
The array can be divided, for example, into the following blocks:

[2, 1, 5, 1, 2, 2, 2], [], [] with a large sum of 15;
[2], [1, 5, 1, 2], [2, 2] with a large sum of 9;
[2, 1, 5], [], [1, 2, 2, 2] with a large sum of 8;
[2, 1], [5, 1], [2, 2, 2] with a large sum of 6.
The goal is to minimize the large sum. In the above example, 6 is the minimal large sum.

Write a function:

func Solution(K int, M int, A []int) int

that, given integers K, M and a non-empty array A consisting of N integers, returns the minimal large sum.

For example, given K = 3, M = 5 and array A such that:

  A[0] = 2
  A[1] = 1
  A[2] = 5
  A[3] = 1
  A[4] = 2
  A[5] = 2
  A[6] = 2
the function should return 6, as explained above.

Write an efficient algorithm for the following assumptions:

N and K are integers within the range [1..100,000];
M is an integer within the range [0..10,000];
each element of array A is an integer within the range [0..M].
*/

func isSplitValid(A []int, largestSum int, K int) bool {
	sum := 0
	count := 0

	for _, v := range A {
		if sum+v > largestSum {
			sum = 0
			count++
		}
		sum += v
	}
	return count < K
}

func binarySearch(K int, A []int, minSum int, maxSum int) int {
	result := maxSum
	for minSum <= maxSum {
		midSum := (minSum + maxSum) / 2
		if isSplitValid(A, midSum, K) {
			maxSum = midSum - 1
			result = midSum
		} else {
			minSum = midSum + 1
		}
	}
	return result
}

func MinMaxDivision(K int, M int, A []int) int {
	// maxSum is a sum of all elements and minSum is the max element
	minSum, maxSum := 0, 0
	for _, v := range A {
		maxSum += v
		if v > minSum {
			minSum = v
		}
	}
	return binarySearch(K, A, minSum, maxSum)
}
