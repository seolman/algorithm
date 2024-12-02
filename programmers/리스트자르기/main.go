package main

import (
	"fmt"
)

func main() {
    n := 3
    slicer := []int{1, 5, 2}
    num_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(solution(n, slicer, num_list))
}

func solution(n int, slicer []int, num_list []int) []int {
    a, b, c := slicer[0], slicer[1], slicer[2]
    switch n {
    case 1:
        return num_list[:b+1]
    case 2:
        return num_list[a:]
    case 3:
        return num_list[a:b+1]
    case 4:
        return num_list[a:b+1:c]
    }
    return []int{}
}

