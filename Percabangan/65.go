package main

import "fmt"

func main() {
	var a, c int
	var b bool
	fmt.Scan(&a, &b)
	if a >= 500000 && b {
		c = a * 20 / 100
	}
	fmt.Print(c)
}
