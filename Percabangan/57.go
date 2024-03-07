package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a%2 == 0 && b%2 == 0 && c%2 == 0 {
		fmt.Print("genap")
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 {
		fmt.Print("ganjil")
	} else {
		fmt.Print("bukan ganjil atau genap")
	}
}
