package main

import "fmt"

func main() {
	var a, b bool
	var i int
	fmt.Scan(&a, &i, &b)
	fmt.Print(a && b && i >= 19)
}
