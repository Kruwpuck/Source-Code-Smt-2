package main

import (
	"fmt"
)

const hargaGaming = 14000
const hargaBiasa = 10000

type Customer struct {
	nama      string
	lamaMain  int
	ruangan   string
	biaya     int
}

func hitungBiaya(customer *Customer) {
	if customer.ruangan == "gaming" {
		customer.biaya = customer.lamaMain * hargaGaming
	} else if customer.ruangan == "biasa" {
		customer.biaya = customer.lamaMain * hargaBiasa
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var totalBiaya int

	for i := 0; i < n; i++ {
		var nama, ruangan string
		var lamaMain int

		fmt.Scan(&nama, &lamaMain, &ruangan)

		customer := Customer{nama: nama, lamaMain: lamaMain, ruangan: ruangan}
		hitungBiaya(&customer)

		fmt.Printf("Total yang harus dibayar %s sebesar %d\n", customer.nama, customer.biaya)
		totalBiaya += customer.biaya
	}

	fmt.Printf("Biaya total sebesar %d\n", totalBiaya)
}