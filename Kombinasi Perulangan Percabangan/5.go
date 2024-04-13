package main

import "fmt"

func main() {
	var b, c, d int
	var a string
	for true {
		fmt.Scan(&a, &b)
		if a == "x" && b == 0 {
			break
		} else if a == "m" {
			c += b
		} else {
			d += b
		}
	}
	fmt.Print(d, c)

}
