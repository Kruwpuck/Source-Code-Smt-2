package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a == "tidak" {
		fmt.Println("berangkat")
	} else {
		fmt.Scan(&b)
		if b == "ya" {
			fmt.Println("berankat")
		} else {
			fmt.Println("diam di rumah")
		}
	}
}
