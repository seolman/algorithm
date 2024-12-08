package math


func Permutation(n, k int) int {
	if k > n {
		return 0
	}
	return Factorial(n) / Factorial(n-k)
}
