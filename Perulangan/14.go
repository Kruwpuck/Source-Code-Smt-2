package main

import "fmt"

func main() {
	var n, m, o, i float64
	fmt.Scan(&n, &m)
	for i = n; i <= m; i++ {
		o += (2.0 / (i))
	}
	fmt.Printf("%.2f", o)
}
