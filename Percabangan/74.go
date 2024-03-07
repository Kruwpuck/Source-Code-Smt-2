package main

import "fmt"

func main() {
	var a, c int
	var b bool
	fmt.Scan(&a, &b, &c)
	fmt.Print(a >= 10 && b && c == 100)
}
