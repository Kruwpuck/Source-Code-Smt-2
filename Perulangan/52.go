package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a >= b && b >= c {
		fmt.Print(b)
	} else if b >= a && a >= c {
		fmt.Print(a)
	} else if a >= c && c >= b {
		fmt.Print(c)
	} else if b >= c && c >= a {
		fmt.Print(c)
	} else if c >= a && a >= b {
		fmt.Print(a)
	} else if c >= b && b >= a {
		fmt.Print(b)
	}
}
