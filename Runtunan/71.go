package main

import "fmt"

func main() {
	var d1, d2, d3, h1, h2 float64
	fmt.Scan(&d1, &d2)
	d3 = d1 + d2
	h1 = d1 / d3
	h2 = d2 / d3
	fmt.Printf("%.2f %.2f\n", h1, h2)
}
