package main

import "fmt"

func main() {
	var a, b, c int
	for true {
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		} else {
			if a%2 == 1 {
				c += b
			}
		}
	}
	fmt.Println(c)
}
