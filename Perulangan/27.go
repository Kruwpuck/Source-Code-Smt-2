package main

import "fmt"

func main() {
	var i, jml, n, x int
	fmt.Scan(&x, &n)
	for i = 1; i <= n; i++ {
		jml += x + (2500 * (i - 1))
	}
	fmt.Println(jml)
}
