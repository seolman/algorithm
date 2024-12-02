package main

import (
	"fmt"
)

func main() {
    arr := []int{1, 2, 3, 100, 99, 98}
    fmt.Println(solution(arr))
}

func solution(arr []int) int {
    count := 0
    changed := true
    for changed {
        changed = false
        for i, num := range arr {
            if num >= 50 && num % 2 == 0 {
                arr[i] = num / 2
                changed = true
            } else if num < 50 && num % 2 != 0 {
                arr[i] = num * 2 + 1
                changed = true
            }
        }
        if !changed {
            break
        }
        fmt.Println(count, arr)
        count++
    }
    return count
}

func Equal(a, b []int) bool {
    if len(a) == len(b) {
        return false
    }
    
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
