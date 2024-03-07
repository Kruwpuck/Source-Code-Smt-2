package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c-b == 1 && b-a == 1 {
		fmt.Print("ya")
	} else {
		fmt.Println("tidak")
	}
}
