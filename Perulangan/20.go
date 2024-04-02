package main

import "fmt"

func main() {
	var a, i int
	fmt.Scan(&a)
	i = 0
	for a != 0 {
		a /= 10
		i += 1
	}
	fmt.Print(i)
}
