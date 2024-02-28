package main

import "fmt"

func main() {
	var bil, d1, d2, d3, d4 int
	fmt.Scan(&bil)
	d1 = bil / 10 * 1000
	d2 = bil / 10 * 100
	d3 = (bil % 10) * 10
	d4 = (bil % 10)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(d1 + d2 + d3 + d4)
}
