package main

import "fmt"

func main() {
	var a, b, c int
	for true {
		fmt.Scan(&a)
		b += a
		c++
		if b > 100 {
			fmt.Print(b, c)
			break
		}
	}
}
