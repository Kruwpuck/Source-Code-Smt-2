package main

import (
	"fmt"
	"math"
)

func main() {
	var R, T, SbX, SbY float64
	fmt.Scan(&R, &T)
	SbX = panjangX(R, T)
	SbY = panjangY(R, T)
	fmt.Printf("%.2f %.2f\n", SbX, SbY)

}
func panjangX(R, T float64) float64 {
	var SbX float64

	SbX = R * math.Cos(T*math.Pi/180)
	return SbX
}

func panjangY(R, T float64) float64 {
	var SbY float64
	SbY = R * math.Sin(T*math.Pi/180)
	return SbY

}
