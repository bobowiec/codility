package greedy

func GreedyCanoeistB(W []int, k int) int {
	canoes := 0
	j := 0
	i := len(W) - 1
	for i >= j {
		if W[i]+W[j] <= k {
			j++
		}
		canoes++
		i--
	}
	return canoes
}
