package fibonacci

func generateFibonacci(N int) ([]int, map[int]bool) {
	fib := []int{0, 1}
	fibSet := map[int]bool{0: true, 1: true}

	for {
		nextFib := fib[len(fib)-1] + fib[len(fib)-2]
		if nextFib > N {
			break
		}
		fib = append(fib, nextFib)
		fibSet[nextFib] = true
	}

	return fib, fibSet
}

func isSumOfTwoFibs(x int, fib []int, fibSet map[int]bool) bool {
	for _, f := range fib {
		if f > x {
			break
		}
		if _, exists := fibSet[x-f]; exists {
			return true
		}
	}
	return false
}

func SumOfFibo(numbers []int, m int) []bool {
	fib, fibSet := generateFibonacci(m)
	result := make([]bool, len(numbers))

	for i, x := range numbers {
		result[i] = isSumOfTwoFibs(x, fib, fibSet)
	}

	return result
}
