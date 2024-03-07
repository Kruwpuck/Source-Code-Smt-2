package main

import "fmt"

func main() {
	var c int
	for i := 1; i <= 99; i += 2 {
		c += i
	}
	fmt.Print(c)
}
