package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	if b%a != 0 {
		c++
	}
	c += (b / a)
	fmt.Print(c)
}
