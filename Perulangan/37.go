package main

import "fmt"

func main() {
	var n, m, jum int
	fmt.Scan(&n, &m)
	for n <= m {
		jum += n
		n++
	}
	fmt.Print(jum)
}
