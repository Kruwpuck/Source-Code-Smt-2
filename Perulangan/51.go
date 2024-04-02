package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	stop := false
	for !stop {

		fmt.Println(a)
		a++
		stop = a > b
	}

}
