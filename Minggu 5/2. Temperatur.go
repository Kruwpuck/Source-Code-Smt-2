package main

import "fmt"

func main() {
	var C, R, F, K float64
	fmt.Scan(&C)
	konversi(C, &R, &F, &K)
	fmt.Println(R, F, K)
}
func konversi(C float64, R, F, K *float64) {
	*R = 4.0 / 5.0 * C
	*F = 9.0/5.0*C + 32.0
	*K = C + 273.15
}
