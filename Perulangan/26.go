package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c)
	d = (b * b) - (4 * a * c)
	if d >= 0.0 {
		fmt.Println("memiliki titik potong dengan sumbu-x")
	} else {
		fmt.Println("tidak memiliki titik potong dengan sumbu-x")
	}
}
