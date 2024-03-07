package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a%2 == 0 && b%2 == 0 {
		fmt.Print(a * b)
	} else if a%2 != 0 && b%2 != 0 {
		fmt.Print(a + b)
	} else {
		fmt.Print(0)
	}
}
