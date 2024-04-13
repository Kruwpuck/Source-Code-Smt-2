package main

import "fmt"

const nMax int = 100

type dataNilai [nMax]int

func main() {
	var arrNilai dataNilai
	var n int
	fmt.Scan(&n)
	isiArray(n, &arrNilai)
}

func isiArray(a int, b *dataNilai) {
	/*I.S. Terdefinisi bilangan bulat dan array dataNilai
	  F.S. Jika n genap maka mengisi arrNilai dan melanjutkan pemeriksaan,
	  jika n ganjil maka menampilkan pesan error*/
	if a%2 == 0 {
		for i := 0; i < a; i++ {
			fmt.Scan(&b[i])
		}
		cekKecurangan(a, b)
	} else {
		fmt.Println("Jumlah array invalid !!!")
	}

}

func cekKecurangan(n int, m dataNilai) {
	/*I.S. terdefinisi bilangan bulat dan array dataNilai
	  F.S. menampilkan avg1,avg2 dan pesan kemungkinan terjadinya kecurangan apabila avg2 - avg1 >= 20 */
	var s1, s2 float64
	for i := 0; i < a; i++ {
		if i < 10 {
			s1 += m[i]
		} else {
			s2 += m[i]
		}
	}

	rata1 := avgSesi1(s1, b)
	rata2 := avgSesi2(s2, b)

	fmt.Println("Rata rata sesi 1 : %.1f", rata1)
	fmt.Println("Rata rata sesi 2 : %.1f", rata2)

	if rata2-rata1 >= 20 {
		fmt.Println("Kemungkinan terjadi kecurangan")
	} else {
		fmt.Println("Tidak terjadi kecurangan")

	}
}

func avgSesi1(a int, b dataNilai) float64 {
	// mengembalikan avg atau rata rata separuh array pertama\
	var c float64
	c = len(b / 2)
	return float64(a) / c
}

func avgSesi2(a int, b dataNilai) float64 {
	// mengembalikan avg atau rata rata separuh array terakhir
	var c float64
	c = len(b / 2)
	return float64(a) / c
}
