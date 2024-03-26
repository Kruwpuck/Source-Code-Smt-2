package main

import "fmt"

func main() {
	var n, i int
	i = 1
	fmt.Scan(&n)
	ganjil(n, i)
}
func ganjil(n, i int) {
	if i == n {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	} else {
		if i%2 != 0 {
			fmt.Print(i, " ")

		}
		ganjil(n, i+1)
	}
}
