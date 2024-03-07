package main

import "fmt"

func main() {
	var c int
	for i := 2; i <= 50; i++ {
		c += i
	}
	fmt.Print(c)
}
