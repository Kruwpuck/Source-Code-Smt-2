package main

import "fmt"

func main() {
	var r, t float64
	fmt.Scan(&r, &t)
	fmt.Print(1.0 / 3.0 * 3.14 * r * r * t)
}
