package main

import "fmt"

func main() {
	var i, a int
	for true {
		fmt.Scan(&a)
		if a != -1 {
			i += a
		} else {
			break
		}
	}
	fmt.Print(i >= 10)
}
