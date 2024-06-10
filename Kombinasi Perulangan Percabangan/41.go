package main

import "fmt"

func main() {
    var n, m, i int
    var total, sign float64
    
    fmt.Scan(&n)
    fmt.Scan(&m)
    total = 0.0
    sign = 1.0
    
    for i = n; i <= m; i++{
        total = total + sign * (3.0 / float64(i))
        sign = sign * -1
    }
    fmt.Printf("%.2f\n", total)
}