package main

import "fmt"

type Mahasiswa struct {
	Nama string
	NIM  int
}

type tabMhs [40]Mahasiswa

func main() {
	var daftarMhs tabMhs
	var input Mahasiswa
	var i int

	fmt.Scan(&input.Nama)
	for input.Nama != "#" && i < 40 {
		fmt.Scan(&input.NIM)
		daftarMhs[i] = input
		i++
		fmt.Scan(&input.Nama, &input.NIM)
	}

	var keluaran string

	fmt.Scan(&keluaran)

	var j int
	for j = 0; j < i; j++ {
		if daftarMhs[j].NIM%2 == 0 && keluaran == "genap" {
			fmt.Print(daftarMhs[j].Nama, " ")

		} else if daftarMhs[j].NIM%2 == 0 && keluaran == "genap" {
			fmt.Print(daftarMhs[j].Nama, " ")
		}
	}
}
