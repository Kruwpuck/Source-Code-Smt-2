package main

import "fmt"

const NMAX int = 50

type tipe_layanan struct {
	kategori string
	harga    int
}

type data_pasien struct {
	nama, rekap 	string
	id  			int
	waktu           periode
	jenis           tipe_layanan
}
type periode struct {
	tanggal, bulan, tahun int
}
type tData_pasien [NMAX]data_pasien
type tLayanan [NMAX]tipe_layanan

func main() {
	var dat_pas tData_pasien
	var dat_lay tLayanan
	dat_lay[0].kategori = "Reguler"
	dat_lay[0].harga = 100000
	dat_lay[1].kategori = "Perunggu"
	dat_lay[1].harga = 200000
	dat_lay[2].kategori = "Emas"
	dat_lay[2].harga = 300000
	dat_pas[0].nama = "susilo"
	dat_pas[0].id = 12
	dat_pas[0].waktu.tahun = 2021
	dat_pas[0].waktu.bulan = 9
	dat_pas[0].waktu.tanggal = 13
	dat_pas[0].jenis = dat_lay[0]
	dat_pas[0].rekap = "buta"
	dat_pas[1].nama = "bambang"
	dat_pas[1].id = 15
	dat_pas[1].waktu.tahun = 2022
	dat_pas[1].waktu.bulan = 12
	dat_pas[1].waktu.tanggal = 25
	dat_pas[1].jenis = dat_lay[1]
	dat_pas[1].rekap = "lumpuh"
	dat_pas[2].nama = "yudhoyono"
	dat_pas[2].id = 17
	dat_pas[2].waktu.tahun = 2023
	dat_pas[2].waktu.bulan = 3
	dat_pas[2].waktu.tanggal = 9
	dat_pas[2].jenis = dat_lay[2]
	dat_pas[2].rekap = "HIV"
	home(dat_pas, dat_lay)
}
func home(A tData_pasien, B tLayanan) {
	var opsi int
	var n,m int = 3,3
	for opsi != 9 {
		fmt.Println("===================================================")
		fmt.Println("Selamat datang Di Layanan Medical Check Up")
		fmt.Println("Pilih Opsi berikut:")
		fmt.Println("1. Penambahan Data Pasien")
		fmt.Println("2. Penghapusan Data Pasien")
		fmt.Println("3. Pengeditan Data Pasien")
		fmt.Println("4. Pencarian Data Pasien")
		fmt.Println("5. Penambahan Paket Layanan")
		fmt.Println("6. Penghapusan Paket Layanan")
		fmt.Println("7. Pengeditan Paket Layanan")
		fmt.Println("8. Menampilkan Data")
		fmt.Println("9. Keluar")
		fmt.Println("===================================================")
		fmt.Scan(&opsi)
		if opsi == 1 {
			main_tambah_pasien(&A, B, &n,m)
		} else if opsi == 2 {
			main_hapus_pasien(A, B, n)
		} else if opsi == 3 {
			main_edit_pasien(A, B, n,m)
		} else if opsi == 4 {
			main_cari_pasien(A, B, n,m)
		} else if opsi == 5 {
			main_tambah_paket(&B,&m)
		} else if opsi == 6 {
			main_hapus_paket(A, B, n,m)
		} else if opsi == 7 {
			main_edit_layanan(A, B, n,m)
		} else if opsi == 8 {
			main_display(*A, *B, n)
		} else if opsi < 1 || opsi > 9 {
			fmt.Println("Opsi Invalid")
		}
	}

}

func main_tambah_pasien(A *tData_pasien, B tLayanan, n *int, m int) {
	var opsi,tahun,bulan,tanggal int
	var id string
	fmt.Println("===================================================")
	fmt.Println("Menu Tambah Pasien")
	fmt.Println("Masukkan Nama Pasien : ")
	fmt.Scan(&A[*n].nama)
	fmt.Println("Masukkan ID Pasien : ")
	fmt.Scan(&id)
	for !cek_id_pasien(*A,*n,id){
		fmt.Println("Masukkan ID Pasien : ")
		fmt.Scan(&id)
	}
	A[*n].id=id
	fmt.Println("Masukkan Rekap Pasien : ")
	fmt.Scan(&A[*n].rekap)
	fmt.Println("Masukkan Waktu Check Up Pasien (YYYY/MM/DD) : ")
	fmt.Scan(&tahun,&bulan,&tanggal)
	for !cek_waktu_pasien(tahun,bulan,tanggal){
		fmt.Println("Masukkan Waktu Check Up Pasien (YYYY/MM/DD) : ")
		fmt.Scan(&tahun,&bulan,&tanggal)
	}
	A[*n].waktu.tahun = tahun
	A[*n].waktu.bulan =  bulan
	A[*n].waktu.tanggal = tanggal
	list_paket(B,m)
	fmt.Println("Masukkan Jenis Paket Pasien Berdasarkan List Diatas : ")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > m{
		fmt.Println("Input Invalid")
		fmt.Println("Masukkan Jenis Paket Pasien Berdasarkan List Diatas : ")
		fmt.Scan(&opsi)
	}
	fmt.Println("===================================================")
	fmt.Println("Data Pasien Berhasil Ditambahkan")
	fmt.Println("===================================================")
	A[*n].jenis = B[opsi-1]
	*n++
}
func list_paket(B tLayanan, m int) {
	fmt.Println("Jenis Paket :")
	for i := 0; i < m; i++ {
		fmt.Printf("%d. %s %d \n", i+1, B[i].kategori, B[i].harga)
	}
}
func main_hapus_pasien(A tData_pasien, B tLayanan, n int) {
	var opsi,y int
	var x string
	var idx int = -1
	fmt.Println("===================================================")
	fmt.Println("Menu Hapus Data Pasien")
	fmt.Println("Cari Data Pasien yang akan Dihapus berdasarkan :")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Print("Pilih Opsi :")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2  {
		fmt.Println("Opsi Invalid ")
		fmt.Print("Pilih Opsi :")
		fmt.Scanln(&opsi)
	}
	if opsi == 1 {
		fmt.Print("Masukkan nama pasien: ")
		fmt.Scan(&x)
		idx = cari_nama(A, n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID pasien: ")
		fmt.Scan(&y)
		idx = cari_id(A, n, y)
	}
	hapus_pasien(&A, B, &n, idx)
	fmt.Println("===================================================")
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
	return idx
}

func cari_id(A tData_pasien, n, y int) int {
	ascend_insertion_sort(&A, n)
	var left, mid, right int
	left = 0
	right = n - 1
	idx := -1
	for left <= right && A[mid].id != y{
		mid = (left + right) / 2
		if y < A[mid].id{
			right = mid - 1
		}else if y > A[mid].id{
			left = mid + 1
		}else if mid == A[mid].id{
			idx = mid
		}
	}
	return idx
}

func ascend_insertion_sort(A *tData_pasien, n int)  {
	for i := 1; i < n; i++ {
		j := i
		for j > 0{
			if A[i].id > A[j].id{
				A[j-1], A[j] = A[j], A[j-1]
			}
			j = j-1
		}
	}
}

func hapus_pasien(A *tData_pasien, B tLayanan, n *int, idx int) {
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	}else{
		display_pasien(*A,idx)
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data Pasien Di Atas Telah Dihapus")
	}

}

func main_edit_pasien(A tData_pasien, B tLayanan, n,m int) {
	var opsi,y int
	var x string
	var idx int = -1
	fmt.Println("===================================================")
	fmt.Println("Menu Edit Data Pasien")
	fmt.Println("Cari Data Pasien yang akan Diedit berdasarkan :")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")
	fmt.Print("Pilih Opsi :")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2  {
		fmt.Println("Opsi Invalid ")
		fmt.Print("Pilih Opsi :")
		fmt.Scanln(&opsi)
	}
	if opsi == 1 {
		fmt.Print("Masukkan nama pasien: ")
		fmt.Scan(&x)
		idx = cari_nama(A, n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID pasien: ")
		fmt.Scan(&y)
		idx = cari_id(A, n, y)
	}
	fmt.Println("===================================================")
	edit_pasien(&A, &B,n,m, idx)
}

func edit_pasien(A *tData_pasien, B *tLayanan,n,m,idx int) {
	var opsi,tahun,bulan,tanggal int
	var id int
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	}else{
		fmt.Println("Data yang akan Diedit")
		fmt.Println("1. Nama")
		fmt.Println("2. ID")
		fmt.Println("3. Waktu Check Up")
		fmt.Println("4. Rekap")
		fmt.Println("5. Paket Layanan")
		fmt.Print("Masukkan Opsi (1/2/3/4/5): ")
		fmt.Scan(&opsi)
		for opsi < 1 || opsi > 5 {
			fmt.Println("Opsi Invalid")
			fmt.Print("Masukkan Opsi (1/2/3/4/5): ")
			fmt.Scan(&opsi)
		}
		if opsi == 1 {
			fmt.Print("Masukkan Nama Pasien yang Baru: ")
			fmt.Scan(&A[idx].nama)
		} else if opsi == 2 {
			fmt.Print("Masukkan ID Pasien yang Baru: ")
			fmt.Scan(&id)
			for !cek_id_pasien(*A,n,id){
				fmt.Print("Masukkan ID Pasien yang Baru: ")
				fmt.Scan(&id)
			}
			A[idx].id = id
		} else if opsi == 3 {
			fmt.Print("Masukkan Waktu Check Up Pasien yang Baru (YYYY/MM/DD): ")
			fmt.Scan(&tahun,&bulan,&tanggal)
			for !cek_waktu_pasien(tahun,bulan,tanggal){
				fmt.Println("Masukkan Waktu Check Up Pasien (YYYY/MM/DD) : ")
				fmt.Scan(&tahun,&bulan,&tanggal)
			}
			A[idx].waktu.tahun = tahun
			A[idx].waktu.bulan =  bulan
			A[idx].waktu.tanggal = tanggal
		} else if opsi == 4 {
			fmt.Print("Masukkan Rekap Pasien yang Baru: ")
			fmt.Scan(&A[idx].rekap)
		} else if opsi == 5 {
			list_paket(*B,m)
			fmt.Print("Masukkan Paket Layanan Pasien yang Baru Berdasarkan List Diatas: ")
			fmt.Scan(&opsi)
			A[idx].jenis = B[opsi-1]
		}
		fmt.Println("===================================================")
		fmt.Println("Data Pasien Telah Diperbaharui")

	}
}

func main_cari_pasien(A tData_pasien, B tLayanan, n,m int) {
	var opsi int
	fmt.Println("Menu Cari Pasien")
	fmt.Println("1. Cari Pasien Individu")
	fmt.Println("2. Cari Pasien Berdasarkan Periode")
	fmt.Println("3. Cari Pasien Berdasarkan Paket")
	fmt.Print("Masukkan Opsi (1/2/3): ")
	fmt.Scan(&opsi)
	for opsi > 3 || opsi < 1 {
		fmt.Println("Opsi Invalid")
		fmt.Print("Masukkan Opsi (1/2/3): ")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		cari_pasien_individu(A, B, n)
	} else if opsi == 2 {
		cari_pasien_periode(A, B, n)
	} else if opsi == 3 {
		cari_pasien_paket(A, B, n,m)
	}
	fmt.Println("===================================================")
}

func cari_pasien_individu(A tData_pasien, B tLayanan, n int) {
	var opsi,y int
	var x string
	fmt.Println("Menu Cari Pasien")
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan ID")
	fmt.Print("Pilih Opsi (1/2): ")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > 2{
		fmt.Println("Input Invalid")
		fmt.Print("Pilih Opsi (1/2): ")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Print("Masukkan Nama Pasien: ")
		fmt.Scan(&x)
		display_pasien(A, n, cari_nama(A, n, x))
	} else if opsi == 2 {
		fmt.Print("Masukkan ID Pasien: ")
		fmt.Scan(&y)
		display_pasien(A, n, cari_id(A, n, y))
	}
}
func cari_pasien_periode(A tData_pasien, B tLayanan, n int) {

	var y1,m1,d1,y2,m2,d2 int
	fmt.Println("Menu Cari Pasien Periodik")
	fmt.Println("Masukkan Waktu Awal (YYYY/MM/DD): ")
	fmt.Scan(&y1,&m1,&d1)
	for !cek_waktu_pasien(y1,m1,d1){
		fmt.Scan(&y1,&m1,&d1)
	}
	fmt.Println("Masukkan Waktu Akhir (YYYY/MM/DD): ")
	fmt.Scan(&y2,&m2,&d2)
	for !cek_waktu_pasien(y2,m2,d2) || y1 < y2 || (y1==y2 && m2 > m1) || (y1==y2 && m2 == m1 && d2 > d1){
		fmt.Scan(&y2,&m2,&d2)
	}
	display_pasien_periodik(A,n,y1,m1,d1,y2,m2,d2)
}
func display_pasien_periodik(A tData_pasien,n, y1,m1,d1,y2,m2,d2 int)  {
	var i int
	var hari1, hari2, hari_cek int
	
	hari1= y1 * 360 + bulan_hari(m1,y1) + d1
	hari2= y2 * 360 + bulan_hari(m2,y2) + d2
	hari_cek = A[i].waktu.tahun * 360 + bulan_hari(A[i].waktu.bulan,A[i].waktu.tahun) + A[i].waktu.tanggal
	for i = 0; i < n; i++ {
		if hari2 >= hari_cek && hari1 <= hari_cek{
			display_pasien(A,i)
		}
	}
}
func bulan_hari(x,y int) int {
	switch x {
	case 1:
		return 0
	case 2:
		return 31
	case 3:
		if cek_tahun_kabisat(y){
			return 60
		}
		return 59
	case 4:
		if cek_tahun_kabisat(y){
			return 91
		}
		return 90
	case 5:
		if cek_tahun_kabisat(y){
			return 121
		}
		return 120
	case 6:
		if cek_tahun_kabisat(y){
			return 152
		}
		return 151
	case 7:
		if cek_tahun_kabisat(y){
			return 182
		}
		return 181
	case 8:
		if cek_tahun_kabisat(y){
			return 213
		}
		return 212
	case 9:
		if cek_tahun_kabisat(y){
			return 244
		}
		return 243
	case 10:
		if cek_tahun_kabisat(y){
			return 274
		}
		return 273
	case 11:
		if cek_tahun_kabisat(y){
			return 305
		}
		return 304
	case 12:
		if cek_tahun_kabisat(y){
			return 335
		}
		return 334
	}
}

func cari_pasien_paket(A tData_pasien, B tLayanan, n,m int) {
	var opsi int
	fmt.Println("Menu Paket Layanan")
	list_paket(B,m)
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&opsi)
	for i := 0; i < n; i++ {
		if A[i].jenis.kategori == B[opsi-1].kategori {
			display_pasien(A, i)
		}
	}

}
func display_pasien(A tData_pasien, n, idx int) {
	fmt.Println("===================================================")
	fmt.Println("Nama Pasien          : ", A[idx].nama)
	fmt.Println("ID Pasien            : ", A[idx].id)
	fmt.Println("Rekap Pasien         : ", A[idx].rekap)
	fmt.Printf("Waktu Check Up Pasien: %d/%d/%d \n", A[idx].waktu.tahun, A[idx].waktu.bulan, A[idx].waktu.tanggal)
	fmt.Println("Jenis Layanan        : ", A[idx].jenis.kategori)
	fmt.Println()
}

func main_edit_layanan(A tData_pasien, B tLayanan, n,m int) {
	var opsi int
	fmt.Println("Opsi Edit Layanan")
	list_paket(*B,m)
	fmt.Print("Pilih Layanan Yang Akan Diedit: ")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > m{
		fmt.Println("Opsi Invalid")
		fmt.Print("Pilih Layanan Yang Akan Diedit: ")
		fmt.Scan(&opsi)
	}
	edit_layanan(A, &B,n,m, opsi-1)

}
func edit_layanan(A tData_pasien, B *tLayanan, n,m, idx int) {
	var nama string
	var harga int
	fmt.Print("Masukkan nama Paket yang Baru: ")
    fmt.Scan(&nama)
	for !cek_nama_paket(*B,m, nama ){
		fmt.Print("Masukkan nama Paket yang Baru: ")
		fmt.Scan(&nama)
	}
	update_nama_layanan(&A,*B,n,idx, nama)
	B[idx].kategori = nama
	fmt.Print("Masukkan harga Paket yang Baru: ")
	fmt.Scan(&harga)
	for harga < 0{
		fmt.Println("Harga Tidak Bisa Negatif")
		fmt.Print("Masukkan harga Paket yang Baru: ")
		fmt.Scan(&harga)
	}
	update_harga_layanan(&A,*B, n,idx)
	B[idx].harga = harga


}
func update_harga_layanan(A *tData_pasien, B tLayanan, n,idx int)  {
	for i := 0; i < n; i++ {
		if A[i].jenis.kategori == B[idx].kategori{
			A[i].jenis = B[idx]
		}
	}
}
func update_nama_layanan(A *tData_pasien, B tLayanan, n,idx int, x string)  {
	for i := 0; i < n; i++ {
		if A[i].jenis.kategori == B[idx].kategori{
			A[i].jenis.kategori = x
		}
	}
}

func main_hapus_paket(A tData_pasien, B tLayanan,n int, m int)  {
	var opsi int
	fmt.Println("Menu Hapus Paket Layanan")
	list_paket(B,m)
	fmt.Print("Masukkan Paket Layanan Yang Ingin Dihapus: ")
	fmt.Scan(&opsi)

	for  opsi < 2 || opsi > m{
		fmt.Println("Input Invalid")
		fmt.Print("Masukkan Paket Layanan Yang Ingin Dihapus: ")
		fmt.Scan(&opsi)
	}
	hapus_pasien_paket(A,B,n,opsi-1)
	fmt.Println("===================================================")
	hapus_paket(&B,&m,opsi-1)
}

func hapus_pasien_paket(A tData_pasien, B tLayanan, n int, idx int)  {
	for i := 0; i < n; i++ {
		if A[i].jenis.kategori == B[idx].kategori{
			hapus_pasien(&A,i,&n)
		}
	}
}


func hapus_paket(B *tLayanan, m *int, idx int)  {
	for i := idx; i < *m-1; i++ {
		B[i] = B[i+1]
	}
	*m--
	fmt.Println("Data Paket Layanan Telah Dihapus")
	fmt.Println("Pasien Yang Menggunakan Paket Layanan Tersebut Diubah ke Paket Reguler")
}

func main_tambah_paket(B *tLayanan, m *int)  {
	var nama string
	fmt.Print("Masukkan Nama Paket: ")
	fmt.Scan(&nama)
	for !cek_nama_paket(*B,*m,nama){
		fmt.Print("Masukkan Nama Paket: ")
		fmt.Scan(&nama)
	}
	B[*m].kategori = nama
	fmt.Print("Masukkan Harga Paket: ")
	fmt.Scan(&B[*m].harga)
	*m++
}
func cek_nama_paket(B tLayanan, m int, x string) bool  {
	var valid bool = true
	for i := 0; i < m; i++ {
		if B[i].kategori == x {
			fmt.Println("Nama Paket Telah Digunakan")
			valid = false
		}
	}
	return valid
}

func main_display(A tData_pasien, B tLayanan, n int) {
	var opsi int
	var y1,m1,d1,y2,m2,d2 int
	fmt.Println("Menu Pemasukkan")
	fmt.Println("1. Periodik Waktu")
	fmt.Println("2. Ascending Waktu")
	fmt.Println("3. Descending Waktu")
	fmt.Println("4. Ascending Paket")
	fmt.Println("5. Descending Paket")
	fmt.Print("Masukkan Opsi (1/2/3/4/5): ")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > 6 {
		fmt.Println("Opsi Invalid")
		fmt.Print("Masukkan Opsi (1/2/3/4/5): ")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Menu Cari Pasien Periodik")
		fmt.Println("Masukkan Waktu Awal (YYYY/MM/DD): ")
		fmt.Scan(&y1,&m1,&d1)
		for !cek_waktu_pasien(y1,m1,d1){
			fmt.Scan(&y1,&m1,&d1)
		}
		fmt.Println("Masukkan Waktu Akhir (YYYY/MM/DD): ")
		fmt.Scan(&y2,&m2,&d2)
		for !cek_waktu_pasien(y2,m2,d2) || y1 < y2 || (y1==y2 && m2 > m1) || (y1==y2 && m2 == m1 && d2 > d1){
			fmt.Scan(&y2,&m2,&d2)
		}
		hitung_pemasukkan(A,n,y1,m1,d1,y2,m2,d2)
	} else if opsi == 2 {
		descending_waktu(A,n)
	} else if opsi == 3 {
		ascending_waktu(A,n)
	} else if opsi == 4 {
		descending_paket(A,n)
	} else if opsi == 5 {
		ascending_paket(A,n)
	}
}

func hitung_pemasukkan(A tData_pasien,n, y1,m1,d1,y2,m2,d2 int) {
	var i,total int
	var hari1, hari2, hari_cek int
	
	hari1= y1 * 360 + bulan_hari(m1,y1) + d1
	hari2= y2 * 360 + bulan_hari(m2,y2) + d2
	hari_cek = A[i].waktu.tahun * 360 + bulan_hari(A[i].waktu.bulan,A[i].waktu.tahun) + A[i].waktu.tanggal
	for i = 0; i < n; i++ {
		if hari2 >= hari_cek && hari1 <= hari_cek{
			total+= A[i].jenis.harga
		}
	}
	fmt.Printf("Total Pemasukkan Mulai Dari %d/%d/%d hingga %d/%d/%d \n adalah sebesar Rp. %d Rupiah", y1,m1,d1,y2,m2,d2,total)
}
func descending_waktu(A tData_pasien, n int) {
	sort_tahun_descend(&A, n)
	sort_bulan_descend(&A, n)
	sort_tanggal_descend(&A, n)
	for i := 0; i < n; i++ {
		display_pasien(A,i)
	}
}

func sort_tahun_descend(A *tData_pasien, n int)  {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu.tahun > A[idx_max].waktu.tahun {
				idx_max = i
			}
		}
		A[pass] , A[idx_max] = A[idx_max] , A[pass]
	}
}

func sort_bulan_descend(A *tData_pasien, n int)  {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu.bulan > A[idx_max].waktu.bulan && A[i].waktu.tahun >= A[idx_max].waktu.tahun {
				idx_max = i
			}
		}
		A[pass] , A[idx_max] = A[idx_max] , A[pass]
	}
}

func sort_tanggal_descend(A *tData_pasien, n int)  {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu.tanggal > A[idx_max].waktu.tanggal && A[i].waktu.tahun >= A[idx_max].waktu.tahun && A[i].waktu.bulan >= A[idx_max].waktu.bulan {
				idx_max = i
			}
		}
		A[pass] , A[idx_max] = A[idx_max] , A[pass]
	}
}

func ascending_waktu(A tData_pasien, n int) {
	sort_tahun_ascend(&A, n)
	sort_bulan_ascend(&A, n)
	sort_tanggal_ascend(&A, n)
	for i := 0; i < n; i++ {
		display_pasien(A,i)
	}
}

func sort_tahun_ascend(A *tData_pasien, n int)  {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu.tahun < A[idx_min].waktu.tahun {
				idx_min = i
			}
		}
		A[pass] , A[idx_min] = A[idx_min] , A[pass]
	}
}

func sort_bulan_ascend(A *tData_pasien, n int)  {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu.bulan < A[idx_min].waktu.bulan && A[i].waktu.tahun <= A[idx_min].waktu.tahun {
				idx_min = i
			}
		}
		A[pass] , A[idx_min] = A[idx_min] , A[pass]
	}
}

func sort_tanggal_ascend(A *tData_pasien, n int)  {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].waktu.tanggal < A[idx_min].waktu.tanggal && A[i].waktu.tahun <= A[idx_min].waktu.tahun && A[i].waktu.bulan <= A[idx_min].waktu.bulan {
				idx_min = i
			}
		}
		A[pass] , A[idx_min] = A[idx_min] , A[pass]
	}
}



func descending_paket(A tData_pasien, n int) {
	var i, pass, idx_max int
	for pass = 0; pass <= n-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].jenis.harga > A[idx_max].jenis.harga {
				idx_max = i
			}
		}
		A[pass] , A[idx_max] = A[idx_max] , A[pass]
	}
	for i := 0; i < n; i++ {
		display_pasien(A,n,i)
	}
}
func ascending_paket(A tData_pasien, n int) {
	var i, pass, idx_min int
	for pass = 0; pass <= n-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= n-1; i++ {
			if A[i].jenis.harga < A[idx_min].jenis.harga {
				idx_min = i
			}
		}
		A[pass] , A[idx_min] = A[idx_min] , A[pass]
	}
	for i := 0; i < n; i++ {
		display_pasien(A,i)
	}
}

func cek_id_pasien(A tData_pasien, n int, x int) bool {
	var i int = 0
	var uniq bool = true
	for i < n {
		if A[i].id == x{
			fmt.Println("ID Telah Digunakan")
			i+=n
			uniq = false
		}else{
			i++
		}
	}
	return uniq
}
func cek_waktu_pasien(x,y,z int) bool {
	var valid bool = true
	var bulan_31 bool = y==1 || y == 3 || y == 5 || y == 7 || y == 8 || y == 10 || y == 12
	if x > 2024 || x < 0{
		fmt.Println("Input Tahun Invalid")
		valid = false
	}
	if y < 1 || y > 12{
		fmt.Println("Input Bulan Invalid")
		fmt.Println("Tidak Ada Bulan Dengan Angka Tersebut")
		valid = false
	}
	if z < 1 || z > 31{
		if !bulan_31 && (z < 1 || z > 30){
			fmt.Println("Input Tanggal Invalid")
		}else if y == 2 && cek_tahun_kabisat(x) && (z < 1 || z > 29) {
			fmt.Println("Input Tanggal Invalid")
		}else if y == 2 && !cek_tahun_kabisat(x) && (z < 1 || z > 28){
			fmt.Println("Input Tanggal Invalid")
		}
		valid = false
	}
	return valid
}
func cek_tahun_kabisat(x int) bool {
    if x%400 == 0 {
        return true
    }
    if x%100 == 0 {
        return false
    }
    if x%4 == 0 {
        return true
    }
    return false
}
