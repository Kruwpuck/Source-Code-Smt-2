const NMax int = 1000

type tabInt [NMax]int

func isiArray(arrInt *tabInt, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat sebanyak n.
	  F.S. Array arrInt berisi data yang diberikan*/
	var masuk, i int //variabel untuk input dan variabel untuk loop
	for i = 0; i < n; i++ {
		fmt.Scan(&masuk)
		for masuk < 0 { // loop untuk mengecek apakah inputan < 0 , jika benar maka meminta inputan ulang
			fmt.Scan(&masuk)
		}
		arrInt[i] = masuk
	}
}

func cekGanjil(x int) bool {
	/*Mengembalikan true apabila x adalah bilangan ganjil dan false jika sebaliknya*/
	return x%2 != 0
}

func prosesAngkaGanjil(arrInt tabInt, n int) {
	/*I.S. Terdefinisi array arrInt, array tidak kosong.
	  F.S. Menampilkan semua bilangan ganjil dalam array dan bilangan ganjil terbesar dalam array arrInt*/
	var i, max int
	max = 0
	for i = 0; i < n; i++ {
		ganjil := cekGanjil(arrInt[i])
		if ganjil {
			fmt.Print(arrInt[i], " ") // menampilkan angka ganjil
			if max < arrInt[i] {      // mencari nilai tertinggi dari kumpulan angka ganjil
				max = arrInt[i]
			}
		}
	}
	if max == 0 { // ketika tidak ada nilai ganjil tertinggi maka yang ditampilkan "-"
		fmt.Print("-")
	}
	fmt.Println("")
	fmt.Println(max) // tampilkan nilai tertinggi dari  kumpulan angka ganjil
}