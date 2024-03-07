package main

import "fmt"

func main() {
	var a, d1, d4 int
	fmt.Scan(&a)
	d1 = a / 1000
	d4 = a % 10
	if d1%2 != 0 && d4%2 == 0 {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}
