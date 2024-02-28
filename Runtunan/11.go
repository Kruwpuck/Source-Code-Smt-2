package main

import "fmt"

func main() {
	var x int
	var y float64
	fmt.Scan(&x)
	y = 1 / float64(x)
	fmt.Print(y)
}
