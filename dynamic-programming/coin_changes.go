package dynamicprogramming

import "math"

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// DynamicCoinChanging calculated min number of coins needed to pay k
// C array of denominations 0 < c0 <= c1 <= ... <= cn-1
func DynamicCoinChanging(C []int, k int) []int {
	n := len(C)
	// create two dimensional array with all zeros
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < C[i-1]; j++ {
			dp[i][j] = dp[i-1][j]
		}
		for j := C[i-1]; j < k+1; j++ {
			dp[i][j] = minimum(dp[i][j-C[i-1]]+1, dp[i-1][j])
		}
	}
	return dp[n]
}

func DynamicCoinChangingOpt(C []int, k int) []int {
	n := len(C)
	dp := make([]int, k+1)
	for i := 1; i < k+1; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i < n+1; i++ {
		for j := C[i-1]; j < k+1; j++ {
			dp[j] = minimum(dp[j-C[i-1]]+1, dp[j])
		}
	}
	return dp
}
