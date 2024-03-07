package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a == "A" && (b == "B") {
		fmt.Println(true)
	} else if a == "E" && b == "C" {
		fmt.Println(true)
	} else if a == "B" && (b == "A" || b == "C" || b == "D") {
		fmt.Println(true)
	} else if a == "C" && (b == "E" || b == "B") {
		fmt.Println(true)
	} else if a == "D" && (b == "B") {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
