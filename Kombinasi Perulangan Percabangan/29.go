package main

import "fmt"

func main() {
    var rounds int
    fmt.Scan(&rounds)

    var totalA, totalB, totalC int
    for i := 0; i < rounds; i++ {
        var a, b, c int
        fmt.Scan(&a, &b, &c)
        totalA += a
        totalB += b
        totalC += c
    }

    winner := 'A'
    minTotal := totalA
    if totalB < minTotal {
        winner = 'B'
        minTotal = totalB
    }
    if totalC < minTotal {
        winner = 'C'
        minTotal = totalC
    }

    fmt.Printf("%c %d\n", winner, minTotal)
}