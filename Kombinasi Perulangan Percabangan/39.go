package main

import "fmt"

func main() {
	var a, b int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a%2 != 0 {
			b++
		}
	}
	if b == 5 {
		fmt.Print("ganjil semua")
	} else if b > 0 {
		fmt.Print("ganjil sebagian")
	} else {
		fmt.Print("tidak ada bilangan ganjil")
	}
}
