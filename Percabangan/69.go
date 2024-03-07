package main

import "fmt"

func main() {
	var a, b, c int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if a%2 == 0 {
			b++
		} else {
			c++
		}
	}
	fmt.Print(b, c)
}
