package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

var a, b titik

func main() {
	var hitungawal float64
	fmt.Scan(&a.x, &a.y, &b.x, &b.y)
	hitungawal = jarak(a, b)
	fmt.Print((hitungawal))
}
func jarak(a, b titik) float64 {
	var hasil float64
	hasil = akar(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
	return hasil
}
func akar(masukkan float64) float64 {
	return math.Sqrt(masukkan)
}
