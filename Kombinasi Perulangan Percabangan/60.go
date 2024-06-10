package main

import "fmt"

func main() {
    var numCustomers int
    fmt.Scanln(&numCustomers)

    var discountCount, cashbackCount int

    for i := 0; i < numCustomers; i++ {
        var totalBelanja int
        fmt.Scanln(&totalBelanja)

        if totalBelanja%400 == 0 || (totalBelanja%100 != 0 && totalBelanja%4 == 0) {
            discountCount++
        }
        if totalBelanja%300 == 0 || (totalBelanja%100 != 0 && totalBelanja%3 == 0) {
            cashbackCount++
        }
    }

    fmt.Println(discountCount, cashbackCount)
}