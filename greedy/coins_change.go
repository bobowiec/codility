package greedy

func GreedyCoinChanging(M []int, k int) []int {
	n := len(M)
	result := []int{}
	for i := n - 1; i >= 0; i-- {
		result = append(result, M[i], k/M[i])
		k %= M[i]
	}
	return result
}
