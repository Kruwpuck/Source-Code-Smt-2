package main

import "fmt"

func main() {
	var n, m, jum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		jum += m
	}
	fmt.Print(jum)
}
