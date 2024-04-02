package main

import "fmt"

func main() {
	var a, b, d1, d2, i int
	fmt.Scan(&a)
	d1 = a / 10
	d2 = a % 10
	for true {
		fmt.Scan(&b)
		i++
		if b == d1 || b == d2 {
			fmt.Print(i)
			break
		}
	}
}
