package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a)
	c = 1
	for i := 0; i < int(a); i++ {
		fmt.Scan(&b)
		c *= b
	}
	fmt.Print(int(c))
}
