package main

import "fmt"

func main() {
    totalOdd := 0
    totalEven := 0

    for {
        var d1, d2, d3 int
        fmt.Scanln(&d1, &d2, &d3)

        sum := d1 + d2 + d3
        if sum > 15 {
            break
        }

        if d1%2 == 0 && d2%2 == 0 && d3%2 == 0 {
            totalEven++
        } else if d1%2 != 0 && d2%2 != 0 && d3%2 != 0 {
            totalOdd++
        }
    }

    fmt.Println(totalOdd+totalEven)
}