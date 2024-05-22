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
			main_tambah_pasien(A, B, &n,m)
		} else if opsi == 2 {
			main_hapus_pasien(A, B, &n)
		} else if opsi == 3 {
			main_edit_pasien(A, B, n,m)
		} else if opsi == 4 {
			main_cari_pasien(*A, *B, n,m)
		} else if opsi == 5 {
			main_tambah_paket(B,&m)
		} else if opsi == 6 {
			main_hapus_paket(A, B, n,&m)
		} else if opsi == 7 {
			main_edit_layanan(A, B, n,m)
		} else if opsi == 8 {
			main_display(*A, *B, n)
		} else if opsi < 1 || opsi > 9 {
			fmt.Println("Opsi Invalid")
		}
	}

}

func main_tambah_pasien(A *tData_pasien, B *tLayanan, n *int, m int) {
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
	list_paket(*B,m)
	fmt.Println("Masukkan Jenis Paket Pasien Berdasarkan List Diatas : ")
	fmt.Scan(&opsi)
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
func main_hapus_pasien(A *tData_pasien, B *tLayanan, n *int) {
	var opsi int
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
		idx = cari_nama(*A, *n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID pasien: ")
		fmt.Scan(&x)
		idx = cari_id(*A, *n, x)
	}
	hapus_pasien(A, B, n, idx)
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

func cari_id(A tData_pasien, n int, x string) int {
	var idx int = -1
	i := 0
	for i < n && idx == -1 {
		if A[i].id == x {
			idx = i
		}
		i++
	}
	return idx
}

func hapus_pasien(A *tData_pasien, B *tLayanan, n *int, idx int) {
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	}else{
		display_pasien(*A,*n,idx)
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Data Pasien Di Atas Telah Dihapus")
	}

}

func main_edit_pasien(A *tData_pasien, B *tLayanan, n,m int) {
	var opsi int
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
		idx = cari_nama(*A, n, x)
	} else if opsi == 2 {
		fmt.Print("Masukkan ID pasien: ")
		fmt.Scan(&x)
		idx = cari_id(*A, n, x)
	}
	fmt.Println("===================================================")
	edit_pasien(A, B,n,m, idx)
}

func edit_pasien(A *tData_pasien, B *tLayanan,n,m,idx int) {
	var opsi,tahun,bulan,tanggal int
	var id string
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
	var opsi int
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
			fmt.Println("===================================================")
			display_pasien(A, n, i)
		}
	}
}
func display_pasien_bulan(A tData_pasien, n, x, y int) {
	for i := 0; i < n; i++ {
		if A[i].waktu.tahun == x && A[i].waktu.bulan == y {
			fmt.Println("===================================================")
			display_pasien(A, n, i)
		}
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
			display_pasien(A, n, i)
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

func main_edit_layanan(A *tData_pasien, B *tLayanan, n,m int) {
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
	edit_layanan(A, B,n,m, opsi-1)

}
func edit_layanan(A *tData_pasien, B *tLayanan, n,m, idx int) {
	var nama string
	if idx == 0{
		fmt.Println("Nama Paket Reguler Tidak Dapat Diedit")
		fmt.Print("Masukkan harga Paket yang Baru: ")
		fmt.Scanln(&B[idx].harga)
		update_harga_layanan(&*A,*B, n,0)
	}else{
		fmt.Print("Masukkan nama Paket yang Baru: ")
		fmt.Scan(&nama)
		for !cek_nama_paket(*B,m, nama ){
			fmt.Print("Masukkan nama Paket yang Baru: ")
			fmt.Scan(&nama)
		}
		update_nama_layanan(&*A,*B,n,idx, nama)
		fmt.Scanln(&B[idx].kategori)
		fmt.Print("Masukkan harga Paket yang Baru: ")
		fmt.Scanln(&B[idx].harga)
		update_harga_layanan(&*A,*B, n,idx)
	}
	fmt.Println("Paket Layanan Telah Diperbarui")
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

func main_hapus_paket(A *tData_pasien, B *tLayanan,n int, m *int)  {
	var opsi int
	fmt.Println("Menu Hapus Paket Layanan")
	list_paket(*B,*m)
	fmt.Print("Masukkan Paket Layanan Yang Ingin Dihapus: ")
	fmt.Scan(&opsi)

	if opsi == 1{
		fmt.Println("Paket Reguler Tidak Dapat Dihapus")
	}else if opsi < 2 || opsi > *m{
		fmt.Println("Paket Layanan Tidak Ada")
		fmt.Print("Masukkan Paket Layanan Yang Ingin Dihapus: ")
		fmt.Scan(&opsi)
	}
	transfer_paket_pasien(&*A, &*B,n,opsi-1)
	fmt.Println("===================================================")
	hapus_paket(&*B,&*m,opsi-1)
}

func transfer_paket_pasien(A *tData_pasien, B *tLayanan, n, idx int)  {
	for i := 0; i < n; i++ {
		if A[i].jenis.kategori ==  B[idx].kategori{
			A[i].jenis = B[0]
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
	fmt.Print("Masukkan Harga Pake: ")
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
		descending_waktu(A,n)
	} else if opsi == 4 {
		ascending_waktu(A,n)
	} else if opsi == 5 {
		descending_paket(A,n)
	} else if opsi == 6 {
		ascending_paket(A,n)
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

func descending_waktu(A tData_pasien, n int) {
	sort_tahun_descend(&A, n)
	sort_bulan_descend(&A, n)
	sort_tanggal_descend(&A, n)
	for i := 0; i < n; i++ {
		display_pasien(A,n,i)
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
		display_pasien(A,n,i)
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
		display_pasien(A,n,i)
	}
}
// func swap(A, B *tData_pasien) {
// 	var c tData_pasien
// 	*A = c
// 	*A = *B
// 	*B = c
// }
func cek_id_pasien(A tData_pasien, n int, x string) bool {
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
