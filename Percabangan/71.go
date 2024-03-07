package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if b == d {
		fmt.Print((a + c))
	} else {
		fmt.Print(0)
	}
}
