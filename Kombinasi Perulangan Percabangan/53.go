package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    sum := 0
    for i := 1; i <= n*2; i += 2 {
        sum += i
    }

    fmt.Println(sum)
}