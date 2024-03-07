package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if math.Pow(a, b) == c {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}
}
