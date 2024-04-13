package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if i%2 == 0 && i > 0 {
		fmt.Print("ya")
	} else {
		fmt.Print("tidak")
	}
}
