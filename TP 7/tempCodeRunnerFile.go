package main

import "fmt"

// Tipe bentukan struct titik dengan field x dan y (float64) dan warna (string)
type titik struct {
	x, y  float64
	warna string
}

func pembuatan_titik_baru(x, y float64, w string) titik {
	// Mengembalikan sebuah titik dengan koordinat x dan y, serta warna w
	return titik{x, y, w}
}

func main() {
	var x1, y1, x2, y2 float64
	var w1, w2 string

	// baca data x1, y1, w1, x2, y2, w2
	fmt.Scan(&x1, &y1, &w1, &x2, &y2, &w2)

	// pembuatan titik t1
	t1 := pembuatan_titik_baru(x1, y1, w1)

	// pembuatan titik t2
	t2 := pembuatan_titik_baru(x2, y2, w2)

	// pencetakan titik t1 dan t2
	fmt.Printf("Data titik 1: Koordinat (%.2f, %.2f), warna %s\n", int(t1.x), int(t1.y), t1.warna)
	fmt.Printf("Data titik 2: Koordinat (%.2f, %.2f), warna %s\n", int(t2.x), int(t2.y), t2.warna)
}
