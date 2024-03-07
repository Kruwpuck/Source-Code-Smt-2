package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a == 1 {
		fmt.Println(b * 2980)
	} else if a == 2 {
		fmt.Println(b * 4500)
	} else if a == 3 {
		fmt.Println(b * 9980)
	} else if a == 4 {
		fmt.Println(b * 4490)
	} else if a == 5 {
		fmt.Println(b * 6870)
	}
}
