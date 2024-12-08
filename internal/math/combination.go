package math

import(
)

func Combination(n, k int) int {
	if k > n {
		return 0
	}
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
