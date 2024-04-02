package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for true {
		fmt.Scan(&b)
		a -= b
		if a <= 0 {
			fmt.Println(true)
			break
		}
		fmt.Println(false)
	}
}
