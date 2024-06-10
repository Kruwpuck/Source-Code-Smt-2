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
    maxTotal := totalA
    if totalB > maxTotal {
        winner = 'B'
        maxTotal = totalB
    }
    if totalC > maxTotal {
        winner = 'C'
        maxTotal = totalC
    }

    fmt.Printf("%c %d\n", winner, maxTotal)
}