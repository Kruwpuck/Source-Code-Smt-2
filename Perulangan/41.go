package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Print(b)
	} else {
		fmt.Print(a)
	}
}
