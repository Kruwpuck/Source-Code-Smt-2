package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b, &c)
		d += b
		e += c
	}
	fmt.Print(d, e)
}
