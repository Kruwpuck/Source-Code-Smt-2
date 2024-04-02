package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	if a < 200 {
		b = a * 12
	} else if a >= 200 && a < 400 {
		b = a * 15
	} else if a >= 400 && a < 600 {
		b = a * 18
	} else if a >= 600 {
		b = a * 20
	}
	if b >= 400 {
		b += b * 20 / 100
	}
	if b < 100 {
		b = 100
	}
	fmt.Println(b)
}
