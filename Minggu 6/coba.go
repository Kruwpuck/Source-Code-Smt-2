package main

import "fmt"

func main() {
	var op, op1, op2 string
	var N int
	ambilPerintah(&op, &op1, &op2)
	N = hasilKomputasi(op, nilaiTeks(op1), nilaiTeks(op2))
	tampil(N)
}
func ambilPerintah(op, op1, op2 *string) {
	fmt.Scan(op, op1, op2)
}

func nilaiTeks(a string) int {
	if a == "NOL" {
		return 0
	} else if a == "SATU" {
		return 1
	} else if a == "DUA" {
		return 2
	} else if a == "TIGA" {
		return 3
	} else if a == "EMPAT" {
		return 4
	} else if a == "LIMA" {
		return 5
	} else if a == "ENAM" {
		return 6
	} else if a == "TUJUH" {
		return 7
	} else if a == "DELAPAN" {
		return 8
	} else {
		return 9
	}

}

func teksnilai(a int) string {
	if a == 0 {
		return "NOL"
	} else if a == 1 {
		return "SATU"
	} else if a == 2 {
		return "DUA"
	} else if a == 3 {
		return "TIGA"
	} else if a == 4 {
		return "EMPAT"
	} else if a == 5 {
		return "LIMA"
	} else if a == 6 {
		return "ENAM"
	} else if a == 7 {
		return "ENAM"
	} else if a == 8 {
		return "DELAPAN"
	} else {
		return "SEMBILAN"
	}
}

func hasilKomputasi(op string, op1, op2 int) int {
	if op == "TAMBAH" {
		return op1 + op2
	} else {
		return op1 - op2
	}
}

func tampil(n int) {
	fmt.Println(teksnilai(n))
}
