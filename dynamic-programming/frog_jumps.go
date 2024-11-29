package dynamicprogramming

func FrogJumps(S []int, k int, q int) int {
	n := len(S)
	dp := make([]int, k+1)
	dp[0] = 1
	for j := 1; j < k+1; j++ {
		for i := 0; i < n; i++ {
			if S[i] <= j {
				dp[j] = (dp[j] + dp[j-S[i]]) % q
			}
		}
	}
	return dp[k]
}
