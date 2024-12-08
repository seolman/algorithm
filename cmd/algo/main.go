package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

    imath "github.com/seolman/algorithm/internal/math"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    parts := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

    scanner.Scan()
    parts = strings.Split(scanner.Text(), " ")
	X, _ := strconv.Atoi(parts[0])
	Y, _ := strconv.Atoi(parts[1])

	pwCnt := 0
	cnt := 0

	for i := 0; i < M; i++ {
        scanner.Scan()
        parts := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(parts[0])
		if a != 0 {
			pwCnt++
		} else {
			cnt++
		}
	}

	N -= pwCnt
	result := 1

	if cnt > 0 {
		result *= imath.Combination(N, cnt) * imath.Factorial(cnt)
	}
	N -= cnt
	if N > 0 {
		result *= imath.Permutation(9-(pwCnt+cnt), N)
	}

	fmt.Println(result*X + ((result-1)/3)*Y)
}
