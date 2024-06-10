package main

import "fmt"

func main() {
    var rounds int
    fmt.Scan(&rounds)

    var totalAScore, totalBScore int
    for i := 0; i < rounds; i++ {
        var a1, a2, a3, b1, b2, b3 int
        fmt.Scan(&a1, &a2, &a3, &b1, &b2, &b3)
        totalAScore += a1 + a2 + a3
        totalBScore += b1 + b2 + b3
    }

    winner := '0'
    if totalAScore > totalBScore {
        winner = 'A'
    } else if totalAScore < totalBScore {
        winner = 'B'
    }

    fmt.Printf("%d %d %c\n", totalAScore, totalBScore, winner)
}