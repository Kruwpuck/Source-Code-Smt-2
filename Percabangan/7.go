package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a == "tinggi" {
		fmt.Println("macet")
	} else {
		fmt.Println("tidak macet")
	}
}
