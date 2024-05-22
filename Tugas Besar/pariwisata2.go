package main

import "fmt"

const NMAX int = 20

type fitur [6]string
type tempat_wisata struct {
	nama, id  string
	jarak, biaya int
	fasilitas [6]string
}

var fasilitas = fitur{"Toilet", "Tempat Makan", "Smoking Area", "Air", "Hewan", "Wahana"}

type tabWisata [NMAX]tempat_wisata
var data tabWisata
var nData int

func main() {
	login()
}

func login() {
	var opsi int
	for opsi != 3 {
		fmt.Println("Selamat Datang Di Aplikasi Pariwisata")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login User")
		fmt.Println("3. Exit")
		fmt.Scanln(&opsi)
		if opsi == 1 {
			home_admin()
		} else if opsi == 2 {
			home_user()
		} else if opsi < 1 || opsi > 3 {
			fmt.Println("Opsi Invalid")
		}
	}
}

func home_admin() {
	var opsi int
	opsi = 0
	for opsi != 4 {
		fmt.Println("Menu Admin")
		fmt.Println("1. Tambah Wisata")
		fmt.Println("2. Edit Wisata")
		fmt.Println("3. Hapus Wisata")
		fmt.Println("4. Exit")
		fmt.Scanln(&opsi)
		if opsi == 1 {
			main_tambah_wisata()
		} else if opsi == 2 {
			main_edit_wisata()
		} else if opsi == 3 {
			main_hapus_wisata()
		} else if opsi < 1 || opsi > 4 {
			fmt.Println("Opsi Invalid")
			fmt.Scanln(&opsi)
		}
	}
}

func main_tambah_wisata() {
	var opsi int
	fmt.Println("Masukkan Nama Wisata")
	fmt.Scanln(&data[nData].nama)
	fmt.Println("Masukkan ID Wisata")
	fmt.Scanln(&data[nData].id)
	fmt.Println("Masukkan Jarak Wisata")
	fmt.Scanln(&data[nData].jarak)
	fmt.Println("Masukkan Biaya Wisata")
	fmt.Scanln(&data[nData].biaya)
	fmt.Println("Masukkan Fasilitas Wisata")
	list_fasilitas()
	for opsi != 7 {
		fmt.Scan(&opsi)
		if opsi >= 1 && opsi <= 6 {
			data[nData].fasilitas[opsi-1] = fasilitas[opsi-1]
		}
	}
	nData++
	return
}

func list_fasilitas() {
	for i := 0; i < 6; i++ {
		fmt.Printf("%d. %s \n", i+1, fasilitas[i])
	}
	fmt.Println("7. Selesai")
}

func main_edit_wisata() {
	var opsi, idx int
	var x string
	fmt.Println("Menu Edit Wisata")
	fmt.Println("1. Cari Wisata Berdasarkan Nama")
	fmt.Println("2. Cari Wisata Berdasarkan ID")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Masukkan Nama Wisata")
		fmt.Scan(&x)
		idx = cari_nama_wisata(x)
		if idx != -1 {
			edit_wisata(idx)
			return
		}
	} else if opsi == 2 {
		fmt.Println("Masukkan ID Wisata")
		fmt.Scan(&x)
		idx = cari_id_wisata(x)
		if idx != -1 {
			edit_wisata(idx)
			return
		}
	}
	return
}

func cari_nama_wisata(x string) int {
	var idx int = -1
	for i := 0; i < nData; i++ {
		if data[i].nama == x {
			idx = i
			i = nData
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}
	return idx
}

func cari_id_wisata(x string) int {
	var idx int = -1
	for i := 0; i < nData; i++ {
		if data[i].id == x {
			idx = i
			i = nData
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}
	return idx
}

func edit_wisata(idx int) {
	var opsi int
	fmt.Println("Masukkan Nama Wisata yang Baru")
	fmt.Scanln(&data[idx].nama)
	fmt.Println("Masukkan ID Wisata yang Baru")
	fmt.Scanln(&data[idx].id)
	fmt.Println("Masukkan Jarak Wisata yang Baru")
	fmt.Scanln(&data[idx].jarak)
	fmt.Println("Masukkan Fasilitas Wisata yang Baru")
	list_fasilitas()
	for opsi != 7 {
		fmt.Scan(&opsi)
		if opsi >= 1 && opsi <= 6 {
			data[idx].fasilitas[opsi-1] = fasilitas[opsi-1]
		}
	}
	fmt.Println("Masukkan Biaya Wisata yang Baru")
	fmt.Scanln(&data[idx].biaya)
}

func main_hapus_wisata() {
	var opsi, idx int
	var x string
	fmt.Println("Menu Hapus Wisata")
	fmt.Println("1. Cari Wisata Berdasarkan Nama")
	fmt.Println("2. Cari Wisata Berdasarkan ID")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Masukkan Nama Wisata")
		fmt.Scan(&x)
		idx = cari_nama_wisata(x)
		if idx != -1 {
			hapus_wisata(idx)
			fmt.Println("Data telah dihapus")
		}
	} else if opsi == 2 {
		fmt.Println("Masukkan ID Wisata")
		fmt.Scan(&x)
		idx = cari_id_wisata(x)
		if idx != -1 {
			hapus_wisata(idx)
			fmt.Println("Data telah dihapus")
		}
	}
}

func hapus_wisata(idx int) {
	for i := idx; i < nData-1; i++ {
		data[i] = data[i+1]
	}
	nData--
}

func home_user() {
	var opsi int
	for opsi != 7 {
		fmt.Println("Menu User")
		fmt.Println("1. List Wisata Berdasarkan Jarak Terdekat")
		fmt.Println("2. List Wisata Berdasarkan Jarak Terjauh")
		fmt.Println("3. List Wisata Berdasarkan Biaya Termurah")
		fmt.Println("4. List Wisata Berdasarkan Biaya Termahal")
		fmt.Println("5. List Wisata Berdasarkan Fasilitas")
		fmt.Println("6. Cari Wisata")
		fmt.Println("7. Exit")
		fmt.Scanln(&opsi)

		if opsi == 1 {
			ascend_jarak()
		} else if opsi == 2 {
			descend_jarak()
		} else if opsi == 3 {
			ascend_biaya()
		} else if opsi == 4 {
			descend_biaya()
		} else if opsi == 5 {
			main_cari_fasilitas()
		} else if opsi == 6 {
			main_cari_wisata()
		} else if opsi < 1 || opsi > 7 {
			fmt.Println("Opsi Invalid")
		}
	}
}

func ascend_jarak() {
	var i, pass, idx_min int
	for pass = 0; pass <= nData-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= nData-1; i++ {
			if data[i].jarak < data[idx_min].jarak {
				idx_min = i
			}
		}
		data[pass], data[idx_min] = data[idx_min], data[pass]
	}
	for i := 0; i < nData; i++ {
		display_wisata(i)
	}
}

func descend_jarak() {
	var i, pass, idx_max int
	for pass = 0; pass <= nData-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= nData-1; i++ {
			if data[i].jarak > data[idx_max].jarak {
				idx_max = i
			}
		}
		data[pass], data[idx_max] = data[idx_max], data[pass]
	}
	for i := 0; i < nData; i++ {
		display_wisata(i)
	}
}

func ascend_biaya() {
	var i, pass, idx_min int
	for pass = 0; pass <= nData-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= nData-1; i++ {
			if data[i].biaya < data[idx_min].biaya {
				idx_min = i
			}
		}
		data[pass], data[idx_min] = data[idx_min], data[pass]
	}
	for i := 0; i < nData; i++ {
		display_wisata(i)
	}
}

func descend_biaya() {
	var i, pass, idx_max int
	for pass = 0; pass <= nData-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= nData-1; i++ {
			if data[i].biaya > data[idx_max].biaya {
				idx_max = i
			}
		}
		data[pass], data[idx_max] = data[idx_max], data[pass]
	}
	for i := 0; i < nData; i++ {
		display_wisata(i)
	}
}

func display_wisata(idx int) {
	fmt.Println("Nama Wisata   : ", data[idx].nama)
	fmt.Println("ID Wisata     : ", data[idx].id)
	fmt.Println("Jarak Wisata  : ", data[idx].jarak)
	fmt.Println("Biaya Wisata  : ", data[idx].biaya)
	fmt.Print("Fasilitas Wisata: ")
	for i := 0; i < 6; i++ {
		if data[idx].fasilitas[i] == fasilitas[i] {
			fmt.Print(data[idx].fasilitas[i], " ")
		}
	}
	fmt.Println()
}

func main_cari_fasilitas() {
	var opsi int
	fmt.Println("Masukkan Fasilitas Wahana Yang Ingin Dicari: ")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d. %s \n", i+1, fasilitas[i])
	}
	for opsi < 1 || opsi > 6 {
		fmt.Scan(&opsi)
	}
	cari_fasilitas(opsi - 1)
}

func cari_fasilitas(x int) {
	var ada int
	for i := 0; i < nData; i++ {
		if data[i].fasilitas[x] == fasilitas[x] {
			display_wisata(i)
			ada++
		}
	}
	if ada == 0 {
		fmt.Println("Wisata Dengan Fasilitas Tersebut Tidak Ditemukan")
	}
}

func main_cari_wisata() {
	var opsi, idx int
	var x string
	fmt.Println("Menu Cari Wisata")
	fmt.Println("1. Cari Wisata Berdasarkan Nama")
	fmt.Println("2. Cari Wisata Berdasarkan ID")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Masukkan Nama Wisata")
		fmt.Scan(&x)
		idx = cari_nama_wisata(x)
		if idx != -1 {
			display_wisata(idx)
		}
	} else if opsi == 2 {
		fmt.Println("Masukkan ID Wisata")
		fmt.Scan(&x)
		idx = cari_id_wisata(x)
		if idx != -1 {
			display_wisata(idx)
		}
	}
}
