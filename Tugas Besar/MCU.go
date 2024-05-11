package main

import "fmt"

const NMAX int = 50

type tipe_layanan struct {
	kategori string
	harga    int
}

type data_pasien struct {
	nama, id, rekap string
	waktu           periode
	jenis           tipe_layanan
}
type periode struct {
	tanggal, bulan, tahun int
}
type tData_pasien [NMAX]data_pasien
type tLayanan [3]tipe_layanan

func main() {
	var dat_pas tData_pasien
	var dat_lay tLayanan
	dat_lay[0].kategori = "Perak"
	dat_lay[0].harga = 100000
	dat_lay[1].kategori = "Emas"
	dat_lay[1].harga = 200000
	dat_lay[2].kategori = "Platinum"
	dat_lay[2].harga = 300000
	dat_pas[0].nama = "susilo"
	dat_pas[0].id = "12"
	dat_pas[0].waktu.tahun = 2021
	dat_pas[0].waktu.bulan = 9
	dat_pas[0].waktu.tanggal = 13
	dat_pas[0].jenis = dat_lay[0]
	dat_pas[0].rekap = "buta"
	dat_pas[1].nama = "bambang"
	dat_pas[1].id = "15"
	dat_pas[1].waktu.tahun = 2022
	dat_pas[1].waktu.bulan = 12
	dat_pas[1].waktu.tanggal = 25
	dat_pas[1].jenis = dat_lay[1]
	dat_pas[1].rekap = "lumpuh"
	dat_pas[2].nama = "yudhoyono"
	dat_pas[2].id = "17"
	dat_pas[2].waktu.tahun = 2023
	dat_pas[2].waktu.bulan = 3
	dat_pas[2].waktu.tanggal = 9
	dat_pas[2].jenis = dat_lay[2]
	dat_pas[2].rekap = "HIV"
	home(&dat_pas, &dat_lay)
}
func home(A *tData_pasien, B *tLayanan) {
	var opsi int
	var n int = 3
	for opsi != 7 {
		fmt.Println("Selamat datang Di Layanan Medical Check Up")
		fmt.Println("Pilih Opsi berikut:")
		fmt.Println("1. Penambahan Data Pasien")
		fmt.Println("2. Penghapusan Data Pasien")
		fmt.Println("3. Pengeditan Data Pasien")
		fmt.Println("4. Pencarian Data Pasien")
		fmt.Println("5. Pengeditan Paket Layanan")
		fmt.Println("6. Menampilkan Data")
		fmt.Println("7. Keluar")

		fmt.Scan(&opsi)
		if opsi == 1 {
			main_tambah_pasien(&*A, &*B, &n)
		} else if opsi == 2 {
			main_hapus_pasien(&*A, &*B, &n)
		} else if opsi == 3 {
			main_edit_pasien(&*A, &*B, &n)
		} else if opsi == 4 {
			main_cari_pasien(*A, *B, n)
		} else if opsi == 5 {
			main_edit_layanan(&*A, &*B, &n)
		} else if opsi == 6 {
			main_display(*A, *B, n)
		} else if opsi < 1 || opsi > 7 {
			fmt.Println("Opsi Invalid")
		}
	}

}

func main_tambah_pasien(A *tData_pasien, B *tLayanan, n *int) {
	var opsi int
	fmt.Println("Menu Tambah Pasien")
	fmt.Println("Masukkan Nama Pasien : ")
	fmt.Scan(A[*n].nama)
	fmt.Println("Masukkan ID Pasien : ")
	fmt.Scan(A[*n].id)
	fmt.Println("Masukkan Rekap Pasien : ")
	fmt.Scan(A[*n].rekap)
	fmt.Println("Masukkan Waktu Check Up Pasien (YYYY/MM/DD) : ")
	fmt.Scan(A[*n].waktu.tahun, A[*n].waktu.bulan, A[*n].waktu.tanggal)
	list_paket(*B)
	fmt.Println("Masukkan Jenis Paket Pasien Berdasarkan List Diatas : ")
	fmt.Scan(&opsi)
	A[*n].jenis = B[opsi-1]
	*n++
}
func list_paket(B tLayanan) {
	fmt.Println("Jenis Paket :")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d. %s %d \n", i, B[i].kategori, B[i].harga)
	}
}
func main_hapus_pasien(A *tData_pasien, B *tLayanan, n *int) {
	var opsi int
	var x string
	var idx int = -1
	fmt.Println("Menu Hapus Pasien")
	fmt.Println("Cari Data Pasien yang akan Dihapus berdasarkan :")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Print("Pilih Opsi :")
	for opsi < 1 && opsi > 2 {
		fmt.Scan(&opsi)
		if opsi == 1 {
			fmt.Print("Masukkan nama pasien: ")
			fmt.Scan(&x)
			idx = cari_nama(*A, *n, x)
		} else if opsi == 2 {
			fmt.Print("Masukkan ID pasien: ")
			fmt.Scan(&x)
			idx = cari_id(*A, *n, x)
		} else {
			fmt.Println("Opsi Invalid ")
		}
	}
	hapus_pasien(&*A, &*B, &*n, idx)
}
func cari_nama(A tData_pasien, n int, x string) int {
	var idx int = -1
	i := 0
	for i < n && idx == -1 {
		if A[i].nama == x {
			idx = i
		}
		i++
	}
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	}
	return idx
}

func cari_id(A tData_pasien, n int, x string) int {
	var idx int = -1
	i := 0
	for i < n && idx == -1 {
		if A[i].id == x {
			idx = i
		}
		i++
	}
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	}
	return idx
}

func hapus_pasien(A *tData_pasien, B *tLayanan, n *int, idx int) {
	i := idx
	for i < *n {
		A[i] = A[i+1]
		i++
	}
	*n--
}

func main_edit_pasien(A *tData_pasien, B *tLayanan, n int) {
	var opsi int
	var x string
	var idx int = -1
	fmt.Println("Menu Edit Pasien")
	fmt.Println("Cari Data Pasien yang akan Diedit berdasarkan :")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Print("Pilih Opsi :")
	for opsi < 1 && opsi > 2 {
		fmt.Scan(&opsi)
		if opsi == 1 {
			fmt.Print("Masukkan nama pasien: ")
			fmt.Scan(&x)
			idx = cari_nama(*A, n, x)
		} else if opsi == 2 {
			fmt.Print("Masukkan ID pasien: ")
			fmt.Scan(&x)
			idx = cari_id(*A, n, x)
		} else {
			fmt.Println("Opsi Invalid ")
		}
	}
	edit_pasien(&*A, &*B, idx)
}

func edit_pasien(A *tData_pasien, B *tLayanan, idx int) {
	var opsi, opsi2 int
	fmt.Println("Data yang akan Diedit")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Println("3. Waktu Check Up")
	fmt.Println("4. Rekap")
	fmt.Println("5. Paket Layanan")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Masukkan Nama Pasien yang Baru: ")
		fmt.Scan(&A[idx].nama)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID Pasien yang Baru: ")
		fmt.Scan(&A[idx].id)
	} else if opsi == 3 {
		fmt.Print("Masukkan Waktu Check Up Pasien yang Baru (YYYY/MM/DD): ")
		fmt.Scan(A[idx].waktu.tahun, A[idx].waktu.bulan, A[idx].waktu.tanggal)
	} else if opsi == 4 {
		fmt.Print("Masukkan Rekap Pasien yang Baru: ")
		fmt.Scan(&A[idx].rekap)
	} else if opsi == 5 {
		list_paket(*B)
		fmt.Print("Masukkan Paket Layanan Pasien yang Baru Berdasarkan List Diatas: ")
		fmt.Scan(&opsi2)
		A[idx].jenis = B[opsi2-1]
	}
}

func main_cari_pasien(A tData_pasien, B tLayanan, n int) {
	var opsi int
	fmt.Println("Menu Cari Pasien")
	fmt.Println("1. Cari Pasien Individu")
	fmt.Println("2. Cari Pasien Berdasarkan Periode")
	fmt.Println("3. Cari Pasien Berdasarkan Paket")
	fmt.Print("Masukkan Opsi (1/2/3): ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		cari_pasien_individu(A, B, n)
	} else if opsi == 2 {
		cari_pasien_periode(A, B, n)
	} else if opsi == 3 {
		cari_pasien_paket(A, B, n)
	}
}

func cari_pasien_individu(A tData_pasien, B tLayanan, n int) {
	var opsi int
	var x string
	fmt.Println("Menu Cari Pasien")
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan ID")
	fmt.Print("Pilih Opsi (1/2): ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Masukkan Nama Pasien: ")
		fmt.Scan(&x)
		display_pasien(A, n, cari_nama(A, n, x))
	} else if opsi == 2 {
		fmt.Print("Masukkan ID Pasien: ")
		fmt.Scan(&x)
		display_pasien(A, n, cari_id(A, n, x))
	}
}
func cari_pasien_periode(A tData_pasien, B tLayanan, n int) {
	var opsi int
	var x, y int
	fmt.Println("Menu Cari Pasien Periodik")
	fmt.Println("1. Cari Tahun Tertentu")
	fmt.Println("2. Cari Bulan Tertentu ")
	fmt.Print("Pilih Opsi (1/2): ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Masukkan Tahun: ")
		fmt.Scan(&x)
		display_pasien_tahun(A,n)
	} else if opsi == 2 {
		fmt.Print("Masukkan Tahun & Bulan (YYYY/MM): ")
		fmt.Scan(&x, &y)
		display_pasien_bulan(A,n)
	}
}
func display_pasien_tahun(A tData_pasien, n)  {
	
}


func cari_pasien_paket(A tData_pasien, B tLayanan, n int) {

}
func display_pasien(A tData_pasien, n, idx int) {
	fmt.Println("Nama Pasien          : ", A[idx].nama)
	fmt.Println("ID Pasien            : ", A[idx].id)
	fmt.Println("Rekap Pasien         : ", A[idx].rekap)
	fmt.Printf("Waktu Check Up Pasien: %d/%d/%d \n", A[idx].waktu.tahun, A[idx].waktu.bulan, A[idx].waktu.tanggal)
	fmt.Println("Jenis Layanan        : ", A[idx].jenis.kategori)
	fmt.Println()
}
func home_edit(A *tData_pasien, N *int) {
	var opsi int
	fmt.Print("Pilih opsi edit: ")
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan ID")
	fmt.Scan(&opsi)
	if opsi == 1 {
		display_pasien(*A, cari_nama(*A, *N))
	} else if opsi == 2 {
		display_pasien(*A, cari_id(*A, *N))
	}
	fmt.Print("Pilih opsi edit: ")
	fmt.Println("1. Edit Nama")
	fmt.Println("2. Edit ID")
	fmt.Println("3. Edit Waktu Masuk")
	fmt.Println("4. Edit Layanan")
	fmt.Scan(&opsi)
	if opsi == 1 {
		display_pasien(*A, cari_nama(*A, *N))
	} else if opsi == 2 {
		display_pasien(*A, cari_id(*A, *N))
	} else if opsi == 3 {

	} else if opsi == 4 {

	}
}
func edit(A *tData_pasien, N, inx int) {

}
func edit_masuk(A *tData_pasien, N *int, idx int) {

}
