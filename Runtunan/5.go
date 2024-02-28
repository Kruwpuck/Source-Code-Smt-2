package main

import (
	"fmt"
	"math"
)

func main() {
	var luas, hasil float64
	fmt.Scan(&luas)
	hasil = math.Sqrt(4.0 * luas / 3.14)
	fmt.Println(hasil)
}
