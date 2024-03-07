package main

import "fmt"

func main() {
	var a, b, c int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if a > 0 {
			b++
		} else if a < 0 {
			c++
		}
	}
	fmt.Print(b, c)
}
