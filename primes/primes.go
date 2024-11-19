package primes

func countDivisors(n int) int {
	i := 1
	counter := 0
	for i*i < n {
		if n%i == 0 {
			counter += 2
		}
		i++
	}

	if i*i == n {
		counter++
	}

	return counter
}

func primality(n int) bool {
	i := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

func coins(n int) int {
	result := 0
	coin := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		k := i
		for k <= n {
			coin[k] = (coin[k] + 1) % 2
			k += i
		}
		result += coin[n]
	}
	return result
}
