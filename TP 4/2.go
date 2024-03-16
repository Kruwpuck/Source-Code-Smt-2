package main

import "fmt"

func main() {
	var pilih int
	stop := false
	for !stop {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			fmt.Println("1")
			hitungJumlah()
		} else if pilih == 2 {
			fmt.Println("1")
			hitungKali()
		} else if pilih == 3 {
			fmt.Println("1")
			hitungBagi()
		} else if pilih == 4 {
			fmt.Println("1")
			stop = true
		}

	}
}
func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")

}
func hitungJumlah() {
	var a, b, c int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan:")
	fmt.Scan(&a, &b)
	c = a + b
	fmt.Println(a, b)
	fmt.Println("Hasil penjumlahan:", c)
}
func hitungKali() {
	var a, b, c int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan:")
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
	c = a * b
	fmt.Println("Hasil perkalian:", c)
}
func hitungBagi() {
	var a, b, c float64
	fmt.Println("Masukkan dua bilangan yang akan dibagikan:")
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
	c = a / b
	if b != 0.0 {
		fmt.Printf("Hasil pembagian: ,%2.f", c)
	}
}
