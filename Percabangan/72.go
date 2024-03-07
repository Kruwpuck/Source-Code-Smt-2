package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b+c == 180 {
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}
}
