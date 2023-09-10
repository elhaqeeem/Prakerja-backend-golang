package main

import (
    "fmt"
    "strings"
)

func count[T any](slice []T, f func(T) bool) int {
    count := 0
    for _, s := range slice {
        if f(s) {
            count++
        }
    }
    return count
}

func main() {
    s := []string{"asd", "qwe", "asd"}
    fmt.Println(
        count(
            s,
            func(x string) bool {
                return strings.Contains(x, "a")
            }),
    )
}
