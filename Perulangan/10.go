package main

import "fmt"

func main() {
	var c int
	for i := 1; i <= 1000; i++ {
		c += i
	}
	fmt.Print(c)
}
