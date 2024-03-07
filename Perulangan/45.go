package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c)
	d = a + b + c
	if a >= 65 && b >= 55 && c >= 50 && d >= 180 {
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}
}
