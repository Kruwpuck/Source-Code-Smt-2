package main

import "fmt"

func main() {
    var x, y, a, b int
    fmt.Scan(&x, &y, &a, &b)

    for i := x; i <= y; i++ {
        if !startsWith(i, a) && !endsWith(i, b) {
            fmt.Println(i)
        }
    }
}

// Fungsi untuk memeriksa apakah suatu bilangan dimulai dengan digit tertentu
func startsWith(n, digit int) bool {
    for n >= 10 {
        n /= 10
    }
    return n == digit
}

// Fungsi untuk memeriksa apakah suatu bilangan diakhiri dengan digit tertentu
func endsWith(n, digit int) bool {
    return n%10 == digit
}