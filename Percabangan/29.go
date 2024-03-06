package main

import "fmt"

func main() {
	var n, m, total int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		total += m
	}
	fmt.Println(total)
}
