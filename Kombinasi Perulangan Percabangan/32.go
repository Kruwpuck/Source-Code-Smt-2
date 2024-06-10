package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var totalPositive, totalNeutral, totalNegative int

    // Baca sentimen dari setiap postingan
    for i := 0; i < n; i++ {
        var positive, neutral, negative int
        fmt.Scan(&positive, &neutral, &negative)

        totalPositive += positive
        totalNeutral += neutral
        totalNegative += negative
    }

    // Tentukan sentimen yang paling banyak
    var resultChar string
    var maxSentiment int

    if totalPositive > totalNeutral && totalPositive > totalNegative {
        resultChar = "+"
        maxSentiment = totalPositive
    } else if totalNeutral > totalPositive && totalNeutral > totalNegative {
        resultChar = "="
        maxSentiment = totalNeutral
    } else {
        resultChar = "-"
        maxSentiment = totalNegative
    }

    fmt.Printf("%s %d\n", resultChar, maxSentiment)
}