package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println((a + 1) * (2*b + 2) * (3*c + 3))
}
