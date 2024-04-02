package main

import "fmt"

func main() {
	var i, jum int
	for true {
		fmt.Scan(&i)
		if i != 0 {
			jum += i
		} else {
			break
		}
	}
	fmt.Print(jum)
}
