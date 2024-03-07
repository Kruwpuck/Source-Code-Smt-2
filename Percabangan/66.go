package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a >= b && a >= c && a >= d {
		fmt.Print(a)
	} else if b >= a && b >= c && b >= d {
		fmt.Print(b)
	} else if c >= b && c >= a && c >= d {
		fmt.Print(c)
	} else if d >= b && d >= c && d >= a {
		fmt.Print(d)
	}
}
