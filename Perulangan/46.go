package main

import "fmt"

func main() {
	var x, d1 int
	fmt.Scan(&x)
	for x != 0 {
		d1 = x % 10
		x /= 10
		fmt.Print(d1, " ")
	}
}
