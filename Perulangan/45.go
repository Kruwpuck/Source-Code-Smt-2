package main

import "fmt"

func main() {
	var a, b string
	var c int
	for true {
		fmt.Scan(&a, &b)
		if a == "admin" && b == "admin" {
			fmt.Println(c, "Login berhasil")
			break
		}
		c++
	}
}
