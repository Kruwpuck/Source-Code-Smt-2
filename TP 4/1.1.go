package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	const pi float64 = 3.14
	*l = pi * r * r
	*k = 2 * pi * r
}
func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}
func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}

func main() {
	var rad, sisi, LL, LP, KL, KP, TL, TP float64
	fmt.Scanln(&rad, &sisi)
	if rad != 0 && sisi != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	}
	for rad != 0 && sisi != 0 {
		hitungLuasKelilingLingkaran(rad, &LL, &KL)
		hitungLuasKelilingPersegi(sisi, &LP, &KP)
		hitungTotal(LL, LP, KL, KP, &TL, &TP)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", rad, sisi, LL, LP, KL, KP, TL, TP)
		fmt.Scan(&rad, &sisi)
	}
}
