package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var hasil bool
	fmt.Scan(&a, &b)
	hasil = math.Sqrt(b) == a
	fmt.Print(hasil)
}
