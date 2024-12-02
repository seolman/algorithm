package main

import(
    "fmt"
    // "testing"
)

func main() {
    fmt.Println(solution("Programmers"))
}

func solution(my_string string) []int {
    m := make([]int, 52)
    for _, str := range my_string {
        if 'A' <= str && str <= 'Z' {
            m[str-'A']++
        } else if 'a' <= str && str <= 'z' {
            m[str-'a'+26]++
        }
    }
    return m
}
