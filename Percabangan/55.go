package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	c = a
	a = b
	b = c
	fmt.Println(a, b)
}
