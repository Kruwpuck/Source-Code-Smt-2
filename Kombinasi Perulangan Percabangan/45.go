package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    count := 0
    for i := 1; i <= 360; i++ {
        if i % x == 0 && i % y != 0 {
            count++
        }
    }

    fmt.Println(count)
}