package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for a <= b {
		fmt.Println(a)
		a++
	}
}
