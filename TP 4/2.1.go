package main
import "fmt"
func main() {
	var pilih int
	var stop bool
	stop = false
	for !stop {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		switch(pilih){
		case 1:
		    fmt.Print("1\n")
			hitungJumlah()
		case 2:
	    	fmt.Print("2\n")
			hitungKali()
		case 3:
		    fmt.Print("3\n")
			hitungBagi()
		}
		 	stop = pilih == 4
		 	if stop{
		 	    fmt.Print("4\n")
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
	fmt.Println("4. Exit ")
	fmt.Println("-----------------------")
}
func hitungJumlah() {
	var a, b,hasil int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	hasil = a + b
	fmt.Print(a," ",b,"\n")
	fmt.Println("Hasil penjumlahan:", hasil)
}
func hitungKali() {
	var a, b, hasil int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	hasil = a * b
	fmt.Print(a," ",b,"\n")
	fmt.Println("Hasil perkalian:", hasil)
}
func hitungBagi() {
	var a, b,hasil float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	fmt.Print(a," ",b,"\n")
	if b != 0 {
		hasil = a / b
		fmt.Println("Hasil pembagian:", hasil)
	} 
}