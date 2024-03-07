package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a <= 10 {
		fmt.Print(a * 2500)
	} else {
		fmt.Print(a * 5000)
	}
}
