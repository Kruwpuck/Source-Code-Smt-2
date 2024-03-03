package main

import "fmt"

func main() {
	var C, K, F, R, i, awal, akhir, gap float64
	fmt.Scan(&awal, &akhir, &gap)
	i = awal
	fmt.Println("    Celcius", "    Reamur", "    Fahrenheit", "   Kelvin")
	for i <= akhir {
		R = reamur(i)
		F = fahrenheit(i)
		K = kelvin(i)
		C = i
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
		i += gap
	}
}
func reamur(i float64) float64 {
	return 4.0 * i / 5.0
}
func kelvin(i float64) float64 {
	return i + 273.0
}
func fahrenheit(i float64) float64 {
	return (9.0 * i / 5.0) + 32.0
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var R, T float64
	fmt.Scan(&R, &T)
	fmt.Println(panjangX(R, T), panjangY(R, T))
}
func panjangX(R, T float64) float64 {
	return R * math.Cos(T*math.Pi)
}
func panjangY(R, T float64) float64 {
	return R * math.Sin(T*math.Pi)
}
