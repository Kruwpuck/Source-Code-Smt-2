package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("menang")
	} else if a < b {
		fmt.Println("kalah")
	} else {
		fmt.Println("draw")
	}
}
