package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for true {
		fmt.Scan(&b)
		if b != 0 {
			a += b
		} else {
			break
		}
	}
	fmt.Print(a)
}
