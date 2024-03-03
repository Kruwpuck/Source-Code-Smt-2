package main

import "fmt"

func main() {
	var gram, kilo, d1, d2, d3 int
	fmt.Scan(&d1, &d2, &d3)
	gram = d1 + d2 + d3
	kilo = gram / 1000
	fmt.Println(kilo, gram%1000)
}
