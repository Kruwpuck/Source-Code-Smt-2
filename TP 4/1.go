package main

import "fmt"

func main() {
	var r, s, lL, lP, kL, kP, totLuas, totKel float64
	fmt.Scan(&r, &s)
	if r != 0 && s != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	}
	for r != 0 && s != 0 {
		hitungLuasKelilingLingkaran(r, &lL, &kL)
		hitungLuasKelilingPersegi(s, &lP, &kP)
		hitungTotal(lL, lP, kL, kP, &totLuas, &totKel)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, lL, lP, kL, kP, totLuas, totKel)
		fmt.Scan(&r, &s)
	}
}
func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	const pi float64 = 3.14
	*l = pi * r * r
	*k = 2.0 * pi * r
}
func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}
func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}
