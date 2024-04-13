package main

import "fmt"

func main() {
	var a string
	var lolos int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a)
		if a == "yes" {
			lolos++
		}
	}
	if lolos >= 2 {
		fmt.Println("lolos")
	} else {
		fmt.Println("tidak lolos")
	}
}
