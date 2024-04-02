package main

import "fmt"

func main() {
	var n, p, l int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p, &l)
		fmt.Println(p*l, 2*(p+l))
	}
}
