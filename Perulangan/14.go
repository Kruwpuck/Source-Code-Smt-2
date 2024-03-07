package main

import "fmt"

func main() {
	var n, m, o, i float64
	fmt.Scan(&n, &m)
	for i = 0.0; i < m; i++ {
		o += (4.0 / (n + i))
	}
	fmt.Printf("%.2f", o)
}
