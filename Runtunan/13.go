package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	fmt.Println((x*x + 1.0/x) * (x*x + 1.0/x))
}
