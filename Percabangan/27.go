package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n%m == 0 {
		fmt.Println("Ya,", n, "kelipatan", m)
	} else {
		fmt.Println("Tidak,", n, "bukan kelipatan", m)
	}
}
