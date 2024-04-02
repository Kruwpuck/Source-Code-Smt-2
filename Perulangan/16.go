package main

import "fmt"

func main() {
	var a int
	var s string
	fmt.Scan(&a, &s)
	for i := 1; i <= a; i++ {
		fmt.Println(s)
	}
}
