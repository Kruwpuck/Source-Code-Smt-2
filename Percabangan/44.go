package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b > 0 && a > 0 {
		fmt.Print(1)
	} else if b < 0 && a < 0 {
		fmt.Print(3)
	} else if b > 0 && a < 0 {
		fmt.Print(2)
	} else {
		fmt.Print(4)
	}
}
