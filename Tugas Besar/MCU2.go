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
			main_tambah_pasien(A, B, &n)
		} else if opsi == 2 {
			main_hapus_pasien(A, B, &n)
		} else if opsi == 3 {
			main_edit_pasien(A, B, n)
		} else if opsi == 4 {
			main_cari_pasien(*A, *B, n)
		} else if opsi == 5 {
			main_edit_layanan(A, B, n)
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
	fmt.Scan(&A[*n].nama)
	fmt.Println("Masukkan ID Pasien : ")
	fmt.Scan(&A[*n].id)
	fmt.Println("Masukkan Rekap Pasien : ")
	fmt.Scan(&A[*n].rekap)
	fmt.Println("Masukkan Waktu Check Up Pasien (YYYY/MM/DD) : ")
	fmt.Scan(&A[*n].waktu.tahun, &A[*n].waktu.bulan, &A[*n].waktu.tanggal)
	list_paket(*B)
	fmt.Println("Masukkan Jenis Paket Pasien Berdasarkan List Diatas : ")
	fmt.Scan(&opsi)
	A[*n].jenis = B[opsi-1]
	*n++
}
func list_paket(B tLayanan) {
	fmt.Println("Jenis Paket :")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d. %s %d \n", i+1, B[i].kategori, B[i].harga)
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
	for opsi < 1 || opsi > 2 {
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
	hapus_pasien(A, B, n, idx)
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
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
		return
	}
	for i := idx; i < *n-1; i++ {
		A[i] = A[i+1]
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
	for opsi < 1 || opsi > 2 {
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
	edit_pasien(A, B, idx)
}

func edit_pasien(A *tData_pasien, B *tLayanan, idx int) {
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
		return
	}
	var opsi int
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
		fmt.Scan(&A[idx].waktu.tahun, &A[idx].waktu.bulan, &A[idx].waktu.tanggal)
	} else if opsi == 4 {
		fmt.Print("Masukkan Rekap Pasien yang Baru: ")
		fmt.Scan(&A[idx].rekap)
	} else if opsi == 5 {
		list_paket(*B)
		fmt.Print("Masukkan Paket Layanan Pasien yang Baru Berdasarkan List Diatas: ")
		fmt.Scan(&opsi)
		A[idx].jenis = B[opsi-1]
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
		display_pasien_tahun(A, n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan Tahun & Bulan (YYYY/MM): ")
		fmt.Scan(&x, &y)
		display_pasien_bulan(A, n, x, y)
	}
}
func display_pasien_tahun(A tData_pasien, n, x int) {
	for i := 0; i < n; i++ {
		if A[i].waktu.tahun == x {
			display_pasien(A, n, i)
		}
	}
}
func display_pasien_bulan(A tData_pasien, n, x, y int) {
	for i := 0; i < n; i++ {
		if A[i].waktu.tahun == x && A[i].waktu.bulan == y {
			display_pasien(A, n, i)
		}
	}
}

func cari_pasien_paket(A tData_pasien, B tLayanan, n int) {
	var opsi int
	fmt.Println("Menu Paket Layanan")
	list_paket(B)
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&opsi)
	for i := 0; i < n; i++ {
		if A[i].jenis.kategori == B[opsi-1].kategori {
			display_pasien(A, n, i)
		}
	}

}
func display_pasien(A tData_pasien, n, idx int) {
	if idx < 0 || idx >= n {
		fmt.Println("Index Out of Range")
		return
	}
	fmt.Println("Nama Pasien          : ", A[idx].nama)
	fmt.Println("ID Pasien            : ", A[idx].id)
	fmt.Println("Rekap Pasien         : ", A[idx].rekap)
	fmt.Printf("Waktu Check Up Pasien: %d/%d/%d \n", A[idx].waktu.tahun, A[idx].waktu.bulan, A[idx].waktu.tanggal)
	fmt.Println("Jenis Layanan        : ", A[idx].jenis.kategori)
	fmt.Println()
}

func main_edit_layanan(A *tData_pasien, B *tLayanan, n int) {
	var opsi int
	fmt.Println("Opsi Edit Layanan")
	list_paket(*B)
	fmt.Print("Pilih Layanan Yang Akan Diedit: ")
	fmt.Scan(&opsi)
	if opsi > 0 && opsi < 4 {
		edit_layanan(A, B, opsi-1)
	} else {
		fmt.Println("Opsi Invalid")
	}
}
func edit_layanan(A *tData_pasien, B *tLayanan, idx int) {
	fmt.Print("Masukkan nama Paket yang Baru: ")
	fmt.Scanln(&B[idx].kategori)
	fmt.Print("Masukkan harga Paket yang Baru: ")
	fmt.Scanln(&B[idx].harga)
	fmt.Println("Paket Layanan Telah Diperbarui")
}

func main_display(A tData_pasien, B tLayanan, n int) {
	var opsi, x, y int
	fmt.Println("Menu Pemasukkan")
	fmt.Println("1. Berdasarkan Tahun")
	fmt.Println("2. Berdasarkan Bulan")
	fmt.Println("3. Ascending Waktu")
	fmt.Println("4. Descending Waktu")
	fmt.Println("5. Ascending Paket")
	fmt.Println("6. Descending Paket")
	fmt.Print("Masukkan Opsi (1/2/3/4/5/6): ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Masukkan Tahun: ")
		fmt.Scan(&x)
		fmt.Printf("Data Pemasukkan Tahun %d adalah sebesar Rp. %d \n", x, hitung_pemasukkan_tahun(A, B, n, x))
	} else if opsi == 2 {
		fmt.Print("Masukkan Tahun & Bulan (YYYY/MM): ")
		fmt.Scan(&x, &y)
		fmt.Printf("Data Pemasukkan Tahun %d Bulan %d adalah sebesar Rp. %d \n", x, y, hitung_pemasukkan_bulan(A, B, n, x, y))
	} else if opsi == 3 {
		descending_waktu()
	} else if opsi == 4 {
		ascending_waktu()
	} else if opsi == 5 {
		descending_paket()
	} else if opsi == 6 {
		ascending_paket()
	}
}
func hitung_pemasukkan_tahun(A tData_pasien, B tLayanan, n, x int) int {
	var total int
	for i := 0; i < n; i++ {
		if A[i].waktu.tahun == x {
			for j := 0; j < 3; j++ {
				if A[i].jenis.kategori == B[j].kategori {
					total += B[j].harga
				}
			}
		}
	}
	return total
}
func hitung_pemasukkan_bulan(A tData_pasien, B tLayanan, n, x, y int) int {
	var total int
	for i := 0; i < n; i++ {
		if A[i].waktu.tahun == x && A[i].waktu.bulan == y {
			for j := 0; j < 3; j++ {
				if A[i].jenis.kategori == B[j].kategori {
					total += B[j].harga
				}
			}
		}
	}
	return total
}

func descending_waktu() {

}
func ascending_waktu() {
	var i, pass, idx int
	for pass = 0; pass <= n-2; pass++ {
		idx = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i] < A[idx] {
				idx = i
			}
		}
	}

}
func descending_paket() {

}
func ascending_paket() {

}
func swap(A, B *int) {
	var c int
	*A = c
	*A = *B
	*B = c
}
