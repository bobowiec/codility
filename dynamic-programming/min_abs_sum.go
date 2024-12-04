package dynamicprogramming

/*
https://app.codility.com/programmers/lessons/17-dynamic_programming/min_abs_sum/

For a given array A of N integers and a sequence S of N integers from the set {−1, 1}, we define val(A, S) as follows:

val(A, S) = |sum{ A[i]*S[i] for i = 0..N−1 }|

(Assume that the sum of zero elements equals zero.)

For a given array A, we are looking for such a sequence S that minimizes val(A,S).

Write a function:

func Solution(A []int) int

that, given an array A of N integers, computes the minimum value of val(A,S) from all possible values of val(A,S) for all possible sequences S of N integers from the set {−1, 1}.

For example, given array:

  A[0] =  1
  A[1] =  5
  A[2] =  2
  A[3] = -2
your function should return 0, since for S = [−1, 1, −1, 1], val(A, S) = 0, which is the minimum possible value.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..20,000];
each element of array A is an integer within the range [−100..100].
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinAbsSum(A []int) int {
	// Step 1: Convert to absolute values and count their frequency
	totalSum := 0
	counts := make(map[int]int)
	for _, a := range A {
		absA := abs(a)
		counts[absA]++
		totalSum += absA
	}

	// Step 2: Use a DP array to calculate the closest sum to totalSum / 2
	dp := make([]int, totalSum+1)
	for i := range dp {
		dp[i] = -1 // -1 indicates that the sum is not reachable
	}
	dp[0] = 0 // Sum 0 is always achievable

	for value, count := range counts {
		for j := 0; j <= totalSum; j++ {
			if dp[j] >= 0 { // If current sum j is achievable
				dp[j] = count
			} else if j >= value && dp[j-value] > 0 {
				dp[j] = dp[j-value] - 1
			}
		}
	}

	// Step 3: Find the closest sum to totalSum / 2
	closestSum := 0
	for i := 0; i <= totalSum/2; i++ {
		if dp[i] >= 0 {
			closestSum = i
		}
	}

	// Step 4: Return the minimal difference
	return abs(totalSum - 2*closestSum)
}
