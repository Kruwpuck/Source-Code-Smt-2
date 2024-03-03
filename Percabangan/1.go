package main

import "fmt"

func main() {
	var a, b, c bool
	fmt.Scan(&a, &b, &c)
	if a == b && b == c {
		fmt.Println("berkemah")
	} else {
		fmt.Println("tidak berkemah")
	}
}
