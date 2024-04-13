package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if b-a == c-b {
		fmt.Print("aritmatika")
	} else if b/a == c/b {
		fmt.Print("geometri")
	} else {
		fmt.Print("bukan aritmatika atau geometri")
	}
}
