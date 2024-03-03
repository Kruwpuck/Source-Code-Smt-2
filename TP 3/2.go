package main

import "fmt"

func main() {
	var C, R, F, K, Cawal, Cakhir, Cstep float64
	fmt.Scanf("%g %g %g %g", &Cawal, &Cakhir, &Cstep)

	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for C = Cawal; C <= Cakhir; C += Cstep {
		R = reamur(C)
		F = fahrenheit(C)
		K = kelvin(C)

		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
	}

}

func reamur(C float64) float64 {
	var reamur float64
	reamur = 4.0 / 5.0 * C
	return reamur
}

func fahrenheit(C float64) float64 {
	var fahrenheit float64
	fahrenheit = (9.0 / 5.0 * C) + 32
	return fahrenheit
}
func kelvin(C float64) float64 {
	var kelvin float64
	kelvin = C + 273
	return kelvin
}
