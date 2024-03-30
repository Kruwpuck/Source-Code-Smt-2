package main

import "fmt"

func persentase(penumpang, kapasitas int) float64 {
	// mengembalikan nilai persentase kapasitas bus yang teris
	return float64(penumpang) / float64(kapasitas) * 100.0
}

func valid(hasil float64) bool {
	// mengembalikan nilai true jika persentase penumpang dari 50% hingga 75%
	// terhadap kapasitas bus
	return hasil >= 50.0 && hasil <= 75.0

}

func main() {
	var kapasitasBus, jumlahPenumpang int
	var cek bool
	var persen float64
	fmt.Scan(&kapasitasBus)
	fmt.Scan(&jumlahPenumpang)
	persen = persentase(jumlahPenumpang, kapasitasBus)
	cek = valid(persen)
	if cek {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}
}
