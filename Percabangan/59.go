package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a <= b && b <= c {
		fmt.Print(a, b, c)
	} else if b <= a && a <= c {
		fmt.Print(b, a, c)
	} else if a <= c && c <= b {
		fmt.Print(a, c, b)
	} else if b <= c && c <= a {
		fmt.Print(b, c, a)
	} else if c <= a && a <= b {
		fmt.Print(c, a, b)
	} else if c <= b && b <= a {
		fmt.Print(c, b, a)
	} else {
		fmt.Print()
	}
}
