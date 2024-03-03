package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var cashb, disc bool
	fmt.Scan(&a, &b, &c, &d, &e)
	cashb = a+b+c == c+d+e
	disc = (b+c+d)%(a+e) == 0
	if cashb {
		fmt.Println("cashback")
	}
	if disc {
		fmt.Println("diskon")
	}
}
