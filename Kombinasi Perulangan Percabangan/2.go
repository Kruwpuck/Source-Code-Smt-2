package main

import "fmt"

func main() {
	var a string
	var b, c int
	fmt.Scan(&a)
	b = len(a) - 1
	for i := 0; i < b; i++ {
		if a[i] == 'a' || a[i] == 'i' || a[i] == 'u' || a[i] == 'e' || a[i] == 'o' {
			c++
		}
	}
	fmt.Print(c)
}
