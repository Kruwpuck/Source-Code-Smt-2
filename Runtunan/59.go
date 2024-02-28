package main

import "fmt"

func main() {
	var bil, d1, d2, d3 int
	fmt.Scan(&bil)
	d1 = bil / 100
	d2 = bil / 10 % 10
	d3 = bil % 10
	fmt.Println(d3*100 + d2*10 + d1)
}
