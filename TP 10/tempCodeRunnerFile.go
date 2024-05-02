package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX]int

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var i int
	i = 0
	var x pemain
	for i < NMAX {
		fmt.Scan(&x.nama)
		if x.nama == "none" {
			i += NMAX
		} else {
			fmt.Scan(&x.nomorPunggung, &x.posisi, &x.tinggi)
			A[i] = x
			i++
			*n++
		}

	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	var i int
	fmt.Println("Data pemain:")
	for i = 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)

	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var max, i int
	var nameMax string
	nameMax = A[0].nama
	max = A[0].tinggi
	for i = 1; i < n; i++ {
		if A[i].tinggi > max {
			nameMax = A[i].nama
			max = A[i].tinggi
		}
	}
	fmt.Println("Pemain dengan badan tertinggi:", nameMax)

}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var min, i int
	var nameMin string
	nameMin = A[0].nama
	min = A[0].tinggi
	for i = 1; i < n; i++ {
		if A[i].tinggi < min {
			nameMin = A[i].nama
			min = A[i].tinggi
		}
	}
	fmt.Println("Pemain dengan badan tertinggi:", nameMin)
}
