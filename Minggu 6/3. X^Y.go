package main

import "fmt"

func main() {
	var n, m, i int
	i = 1
	fmt.Scan(&n, &m)
	fmt.Println(pangkat(n, m, i))

}
func pangkat(n, m, i int) int {
	if i == m {
		return n
	} else {
		return n * pangkat(n, m, i+1)
	}

}
