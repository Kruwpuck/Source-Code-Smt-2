package main

import "fmt"

func main() {
	var s string
	var h bool
	fmt.Scan(&s)
	h = s == "D"
	fmt.Print(h)
}
