package main

import "fmt"

func main() {
	var n int
	var k bool
	fmt.Scan(&n, &k)
	if n >= 17 && k {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
