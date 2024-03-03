package main

import "fmt"

func main() {
	var w, x, y, z int
	fmt.Scan(&x, &y, &z)
	w = y
	y = x
	x = z
	z = w
	fmt.Println(x, y, z)
}
