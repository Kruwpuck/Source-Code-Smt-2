package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a == "ya" && b == "ya" || a == "tidak" {
		fmt.Println("berangkat")
	} else {
		fmt.Println("diam di rumah")
	}
}
