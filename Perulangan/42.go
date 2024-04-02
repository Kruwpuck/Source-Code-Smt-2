package main

import "fmt"

func main() {
	var a, b, c float64
	for true {
		fmt.Scan(&a, &b, &c)
		if a < 0 || b < 0 || c < 0 {
			break
		} else {
			fmt.Println(a >= 50.0 && b >= 50.0 && c >= 50.0)
		}
	}
}
