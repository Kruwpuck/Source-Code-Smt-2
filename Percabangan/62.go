package main

import "fmt"

func main() {
	var bil1, d12, d34, d5 int
	fmt.Scan(&bil1, &d5)
	d12 = (bil1 / 100) * 1000
	d34 = bil1 % 100
	d5 *= 100
	fmt.Println(d12 + d5 + d34)
}
