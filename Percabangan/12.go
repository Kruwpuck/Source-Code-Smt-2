package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, z float64
	var jenis string
	fmt.Scan(&jenis, &x, &y)
	z = math.Pow(x, 2) + math.Pow(y, 2)
	if jenis != "tempur" {
		fmt.Println("tidak tembak")
	} else {
		if math.Sqrt(z) <= 5.0 {
			fmt.Println("tembak")
		} else {
			fmt.Println("tidak tembak")
		}
	}
}
