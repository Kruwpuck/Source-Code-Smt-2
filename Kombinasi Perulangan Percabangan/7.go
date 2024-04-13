package main

import "fmt"

func main() {
	var a string
	var b, c, d int
	for true {
		fmt.Scan(&a, &b)
		if b == 0 {
			break
		} else if a == "bagus" {
			c += b
		} else {
			d += b
		}
	}
	fmt.Print(c, d)
}
