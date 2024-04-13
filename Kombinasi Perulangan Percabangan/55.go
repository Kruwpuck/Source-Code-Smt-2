package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if b-a == c-b {
		fmt.Print("ya")
	} else {
		fmt.Print("tidak")
	}
}
