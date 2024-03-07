package main

import "fmt"

func main() {
	var a int
	var b bool
	fmt.Scan(&a, &b)
	fmt.Print(a >= 10 && b)
}
