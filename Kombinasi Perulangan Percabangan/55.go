package main

import "fmt"

func main() {
    var P, S, X, Y, B int

    fmt.Scan(&P, &S, &X, &Y, &B)

    totalPushUp := 0
    totalSitUp := 0

    for i := 1; i <= B*30; i++ {
        if i%2 == 1 && i%X != 0 && i%Y != 0 {
            totalPushUp += P
        } else if i%2 == 0 && i%X != 0 && i%Y != 0 {
            totalSitUp += S
        }
    }

    fmt.Print(totalPushUp, totalSitUp)
    
}