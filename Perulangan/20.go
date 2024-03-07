package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("terbuka ke atas")
	} else {
		fmt.Println("terbuka ke bawah")

	}
}
