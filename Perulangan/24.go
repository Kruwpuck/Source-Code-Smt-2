package main

import "fmt"

func main() {
	var i, n, x int
	var mean, jml float64
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&x)
		jml += 1 / (float64(x))
	}
	mean = float64(n) / jml
	fmt.Printf("%.2f\n", mean)
}
