package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX]pemain

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
	var pemainBaru pemain
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&pemainBaru.nama)
		if pemainBaru.nama == "none" || *n == NMAX {
			break
		}
		fmt.Scan(&pemainBaru.nomorPunggung, &pemainBaru.posisi, &pemainBaru.tinggi)
		A[i] = pemainBaru
		*n++
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

		<namaN> <nomorPunggungN> <posisiN> <tinggiN>"
	*/
	fmt.Println("Data Pemain:")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var tertinggi int = A[0].tinggi
	var namaTertinggi string = A[0].nama
	for i := 1; i < n; i++ {
		if A[i].tinggi > tertinggi {
			tertinggi = A[i].tinggi
			namaTertinggi = A[i].nama
		}
	}
	fmt.Println("Pemain dengan badan tertinggi:", namaTertinggi)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var terendah int = A[0].tinggi
	var namaTerendah string = A[0].nama
	for i := 1; i < n; i++ {
		if A[i].tinggi < terendah {
			terendah = A[i].tinggi
			namaTerendah = A[i].nama
		}
	}
	fmt.Println("Pemain dengan badan terendah:", namaTerendah)
}
