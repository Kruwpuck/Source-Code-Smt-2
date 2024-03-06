package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c)
	d = (b * b) - (4 * a * c)
	if d >= 0.0 {
		fmt.Println("memiliki akar real")
	} else {
		fmt.Println("tidak memiliki akar real")
	}
}
