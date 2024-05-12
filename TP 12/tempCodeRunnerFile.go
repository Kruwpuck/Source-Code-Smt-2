package main

import "fmt"

func main() {
	var n int
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			n++
		}
	}
	fmt.Print(n)
}
