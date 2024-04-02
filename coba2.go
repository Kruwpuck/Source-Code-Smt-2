package main

import "fmt"

type titik struct {
	x, y float64
}
type lingkaran struct {
	pusat                  titik
	radius, luas, keliling float64
}

func main() {
	var t1, t2 lingkaran

	fmt.Scan(&t1.pusat.x, &t1.pusat.y, &t1.radius)
	fmt.Scan(&t2.pusat.x, &t2.pusat.y, &t2.radius)
	lukel(&t1)
	lukel(&t2)
	cetak(t1)
	cetak(t2)
}
func lukel(k *lingkaran) {
	k.luas = 3.14 * (k.radius * k.radius)
	k.keliling = 2.0 * 3.14 * k.radius
}
func cetak(k lingkaran) {
	fmt.Printf("Lingkaran berpusat di titik (%v, %v) dan berradius %.2f memiliki luas %.2f dan keliling sebesar %.2f\n", k.pusat.x, k.pusat.y, k.radius, k.luas, k.keliling)
}
