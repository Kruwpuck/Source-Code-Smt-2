package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		fmt.Println(p*p, 4*(p))
	}
}
