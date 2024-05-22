package main

import "fmt"

const MAXDATA int = 1000

const MAXPARKIRMOBIL int = 400

const MAXPARKIRMOTOR int = 100

const MAXPETUGAS int = 5

type tPetugas struct {
	id, nama, username, password string
}

type tKendaraan struct {
	tipe, noPol, hari string
	waktu             tWaktuParkir
	vip, valet        bool
	hargaParkir       int
}

type tWaktuParkir struct {
	jamMasuk, jamKeluar, menitMasuk, menitKeluar int
}

type tabPetugas [MAXPETUGAS]tPetugas

type tabDataKendaraan [MAXDATA]tKendaraan

type tabParkirMobil [MAXPARKIRMOBIL]tKendaraan

type tabParkirMotor [MAXPARKIRMOTOR]tKendaraan

func main() {
	var inputMenu, nPetugas int
	var petugas tabPetugas
	for inputMenu != 3 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                  *SELAMAT DATANG*                                  *||*")
		fmt.Println("*||*                                    DI PARKIRAN                                     *||*")
		fmt.Println("*||*                                  *PARIS VAN JOWO*                                  *||*")
		fmt.Println("*||*                         GENA DARMA    // FACHRI MUTTHAWWA                          *||*")
		fmt.Println("*||*                       (103032330095)  //  (103032330095)                           *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*** Menu Utama ***")
		fmt.Println("1. Admin")
		fmt.Println("2. Login Petugas Parkir")
		fmt.Println("3. Keluar")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3) : ")
		fmt.Scan(&inputMenu)
		if inputMenu == 1 {
			adminMenu(&petugas, &nPetugas)
		} else if inputMenu == 2 {
			loginMenu(petugas, nPetugas)
		}
	}
}

func adminMenu(petugas *tabPetugas, nPetugas *int) {
	var input int
	for input != 5 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                    *MENU ADMIN*                                    *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Tambah Data Petugas Parkir")
		fmt.Println("2. Edit Data Petugas Parkir")
		fmt.Println("3. Hapus Data Petugas Parkir")
		fmt.Println("4. Cetak Data Petugas Parkir")
		fmt.Println("5. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5) : ")
		fmt.Scan(&input)
		if input == 1 {
			tambahPetugas(&*petugas, &*nPetugas)
		} else if input == 2 {
			editPetugas(&*petugas, *nPetugas)
		} else if input == 3 {
			hapusPetugas(&*petugas, &*nPetugas)
		} else if input == 4 {
			cetakPetugas(*petugas, *nPetugas)
		}
	}
}

func tambahPetugas(petugas *tabPetugas, nPetugas *int) {
	var i, n int
	var id, usn string
	fmt.Println("============================================================================================")
	fmt.Println("*||*                                *MENU TAMBAH PETUGAS*                               *||*")
	fmt.Println("============================================================================================")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan Banyaknya Petugas Baru : ")
	fmt.Scan(&n)
	if n > MAXPETUGAS-*nPetugas {
		fmt.Println("Banyaknya petugas melebihi batasan maksimum!")
		fmt.Printf("Banyaknya petugas baru yang ditambahkan berubah dari %d menjadi %d\n", n, MAXPETUGAS-*nPetugas)
		n = MAXPETUGAS - *nPetugas
	}
	i = *nPetugas
	for i < n+*nPetugas && i < MAXPETUGAS { //ketimpa euy kalo mau nambah baru pas udah pernah ngisi data
		fmt.Printf("Masukkan ID Petugas %d : ", i+1)
		fmt.Scan(&id)
		for !validId(*petugas, i, id) {
			fmt.Println("ID Petugas tidak valid!")
			fmt.Printf("Masukkan ID Petugas %d : ", i+1)
			fmt.Scan(&id)
		}
		petugas[i].id = id
		fmt.Printf("Masukkan Nama Petugas %d : ", i+1)
		fmt.Scan(&petugas[i].nama)
		fmt.Printf("Masukkan Username Petugas %d : ", i+1)
		fmt.Scan(&usn)
		for !validUsn(*petugas, i, usn) {
			fmt.Println("Username telah digunakan!")
			fmt.Printf("Masukkan Username Petugas %d : ", i+1)
			fmt.Scan(&usn)
		}
		petugas[i].username = usn
		fmt.Printf("Masukkan Password Petugas %d : ", i+1)
		fmt.Scan(&petugas[i].password)
		fmt.Println("Petugas berhasil ditambahkan :)")
		fmt.Println("--------------------------------------------------------------------------------------------")
		i++
	}
	*nPetugas = i
}

func editPetugas(petugas *tabPetugas, nPetugas int) {
	var input int
	for input != 5 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                *MENU EDIT PETUGAS*                                 *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Data Petugas yang ingin diedit :")
		fmt.Println("1. ID Petugas Parkir")
		fmt.Println("2. Nama Petugas Parkir")
		fmt.Println("3. Username Petugas Parkir")
		fmt.Println("4. Password Petugas Parkir")
		fmt.Println("5. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5) : ")
		fmt.Scan(&input)
		if input == 1 || input == 2 || input == 3 || input == 4 {
			editDataPetugas(&*petugas, nPetugas, input)
		}
	}

}

func hapusPetugas(petugas *tabPetugas, nPetugas *int) {
	var idHapus string
	var idx, i int
	fmt.Println("============================================================================================")
	fmt.Println("*||*                                *MENU HAPUS PETUGAS*                                *||*")
	fmt.Println("============================================================================================")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan ID Petugas yang ingin dihapus : ")
	fmt.Scan(&idHapus)
	idx = searchID(*petugas, *nPetugas, idHapus)
	if idx == -1 {
		fmt.Println("ID tersebut tidak terdapat di dalam data Petugas!")
	} else {
		i = idx
		for i < *nPetugas-1 {
			petugas[i] = petugas[i+1]
			i++
		}
		*nPetugas--
		fmt.Printf("Semua data Petugas %d berhasil dihapus :)\n", idx+1)
	}

}

func cetakPetugas(petugas tabPetugas, nPetugas int) { //kali aja mau nambahin urut berdasarkan id / nama
	var i int
	fmt.Println("Data Petugas Tiket Parkir : ")
	for i = 0; i < nPetugas; i++ {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("ID Petugas %d : %s\n", i+1, petugas[i].id)
		fmt.Printf("Nama Petugas %d : %s\n", i+1, petugas[i].nama)
		fmt.Printf("Username Petugas %d : %s\n", i+1, petugas[i].username)
		fmt.Printf("Password Petugas %d : %s\n", i+1, petugas[i].password)
		fmt.Println("--------------------------------------------------------------------------------------------")
	}
}

func validId(petugas tabPetugas, nPetugas int, id string) bool {
	var i int
	var ketemu bool = false
	for i < nPetugas && !ketemu {
		ketemu = petugas[i].id == id
		i++
	}
	return !ketemu
}

func validUsn(petugas tabPetugas, nPetugas int, usn string) bool {
	var i int
	var ketemu bool = false
	for i < nPetugas && !ketemu {
		ketemu = petugas[i].username == usn
		i++
	}
	return !ketemu
}

func searchID(petugas tabPetugas, nPetugas int, id string) int {
	var i, idx int
	idx = -1
	for i < nPetugas && idx == -1 {
		if petugas[i].id == id {
			idx = i
		}
		i++
	}
	return idx
}

func editDataPetugas(petugas *tabPetugas, nPetugas, input int) {
	var idx int
	var id, usn, idEdit string
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan ID Petugas yang ingin diedit : ")
	fmt.Scan(&idEdit)
	idx = searchID(*petugas, nPetugas, idEdit)
	if idx == -1 {
		fmt.Println("ID tersebut tidak terdapat di dalam data Petugas!")
	} else {
		if input == 1 {
			fmt.Printf("Masukkan ID Petugas %d yang baru : ", idx+1)
			fmt.Scan(&id)
			for !validId(*petugas, nPetugas, id) {
				fmt.Println("ID Petugas tidak valid!")
				fmt.Printf("Masukkan ID Petugas %d yang baru : ", idx+1)
				fmt.Scan(&id)
			}
			petugas[idx].id = id
			fmt.Println("ID Petugas berhasil diedit :)")
		} else if input == 2 {
			fmt.Printf("Masukkan Nama Petugas %d yang baru : ", idx+1)
			fmt.Scan(&petugas[idx].nama)
			fmt.Println("Nama Petugas berhasil diedit :)")
		} else if input == 3 {
			fmt.Printf("Masukkan Username Petugas %d yang baru : ", idx+1)
			fmt.Scan(&usn)
			for !validUsn(*petugas, nPetugas, usn) {
				fmt.Println("Username telah digunakan!")
				fmt.Printf("Masukkan Username Petugas %d yang baru : ", idx+1)
				fmt.Scan(&usn)
			}
			petugas[idx].username = usn
			fmt.Println("Username Petugas berhasil diedit :)")
		} else if input == 4 {
			fmt.Printf("Masukkan Password Petugas %d yang baru : ", idx+1)
			fmt.Scan(&petugas[idx].password)
			fmt.Println("Password Petugas berhasil diedit :)")
		}
	}
}

func loginMenu(petugas tabPetugas, nPetugas int) {
	var input, idx int
	var username, password string
	var statusLogin bool
	fmt.Println("============================================================================================")
	fmt.Println("*||*                           *MENU LOG IN PETUGAS PARKIR*                             *||*")
	fmt.Println("============================================================================================")
	for input != 2 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Log In")
		fmt.Println("2. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2) : ")
		fmt.Scan(&input)
		if input == 1 {
			statusLogin = false
			for !statusLogin {
				fmt.Print("Masukkan Username : ")
				fmt.Scan(&username)
				fmt.Print("Masukkan Password : ")
				fmt.Scan(&password)
				idx = searchIdxUsn(petugas, nPetugas, username)
				if cekLogin(petugas, nPetugas, username, password) {
					fmt.Println("Log In berhasil! Menuju ke menu petugas!")
					fmt.Println("--------------------------------------------------------------------------------------------")
					petugasMenu(petugas, idx) // buat setelah login cuy
					statusLogin = true
				} else {
					fmt.Println("Username atau Password Anda Salah! Silahkan input ulang!")
					fmt.Println("--------------------------------------------------------------------------------------------")
				}
			}
		}
	}
}

func cekLogin(petugas tabPetugas, nPetugas int, username, password string) bool {
	var i int
	var ketemu bool = false
	for i < nPetugas && !ketemu {
		ketemu = petugas[i].username == username && petugas[i].password == password
		i++
	}
	return ketemu
}

func searchIdxUsn(petugas tabPetugas, nPetugas int, usn string) int {
	var idx, i int
	idx = -1
	for i < nPetugas && idx == -1 {
		if petugas[i].username == usn {
			idx = i
		}
		i++
	}
	return idx
}

func petugasMenu(petugas tabPetugas, idx int) {
	var input, nMobil, nMotor, nData, totalPendapatan int
	var pMobil tabParkirMobil
	var pMotor tabParkirMotor
	var dataKendaraan tabDataKendaraan
	for input != 6 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                  *MENU PETUGAS*                                    *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("*Selamat Datang Petugas %s, silahkan pilih perintah berikut. Selamat bekerja!*\n", petugas[idx].nama)
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Tambah Data Kendaraan Parkir")
		fmt.Println("2. Edit Data Kendaraan Parkir")
		fmt.Println("3. Hapus Data Kendaraan Parkir")
		fmt.Println("4. Cari Data Kendaraan Parkir")
		fmt.Println("5. Cetak Data Kendaraan Parkir")
		fmt.Println("6. Log Out")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5/6) : ")
		fmt.Scan(&input)
		if input == 1 {
			tambahKendaraan(&pMobil, &pMotor, &dataKendaraan, &nMobil, &nMotor, &nData, &totalPendapatan) // 7.00 - 22.00
		} else if input == 2 {
			editKendaraan(&pMobil, &pMotor, &dataKendaraan, &totalPendapatan, nMobil, nMotor, nData)
		} else if input == 3 {
			hapusDataKendaraan(&pMobil, &pMotor, &dataKendaraan, &nMobil, &nMotor, &nData, &totalPendapatan)
		} else if input == 4 {
			cariKendaraan(dataKendaraan, nData)
		} else if input == 5 {
			cetakKendaraan(pMobil, pMotor, dataKendaraan, nMobil, nMotor, nData)
		}
	}
}

func tambahKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int) {
	var input int
	for input != 3 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                       *MENU TAMBAH KENDARAAN PARKIR*                               *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Kendaraan Masuk Parkir")
		fmt.Println("2. Kendaraan Keluar Parkir")
		fmt.Println("3. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3) : ")
		fmt.Scan(&input)
		if input == 1 {
			masukParkir(&*pMobil, &*pMotor, &*dataKendaraan, &*nMobil, &*nMotor, &*nData)
		} else if input == 2 {
			keluarParkir(&*pMobil, &*pMotor, &*dataKendaraan, &*nMobil, &*nMotor, &*nData, &*totalPendapatan)
		}
	}
}

func masukParkir(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData *int) {
	var input, jamMasuk, menitMasuk int
	var nopol string
	for input != 3 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Jenis Kendaraan :")
		fmt.Println("1. Mobil")
		fmt.Println("2. Motor")
		fmt.Println("3. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3) : ")
		fmt.Scan(&input)
		if input == 1 {
			if *nMobil > MAXPARKIRMOBIL-1 {
				fmt.Println("Parkiran mobil sudah penuh!")
				*nMobil = MAXPARKIRMOBIL - 1
			} else {
				pMobil[*nMobil].tipe = "Mobil"
				fmt.Print("Hari (Minggu,..,Jumat,Sabtu) : ")
				fmt.Scan(&pMobil[*nMobil].hari)
				fmt.Print("Nomor Polisi : ")
				fmt.Scan(&nopol) // nopol unik brow
				for !validNoPolMobil(*pMobil, *nMobil, nopol) || !validNoPolMotor(*pMotor, *nMotor, nopol) {
					fmt.Println("Nomor Polisi tidak valid!")
					fmt.Print("Nomor Polisi :")
					fmt.Scan(&nopol)
				}
				pMobil[*nMobil].noPol = nopol
				fmt.Print("Vip : ")
				fmt.Scan(&pMobil[*nMobil].vip)
				fmt.Print("Valet : ")
				fmt.Scan(&pMobil[*nMobil].valet)
				fmt.Print("Jam Masuk : ")
				fmt.Scan(&jamMasuk, &menitMasuk)
				for jamMasuk < 7 || jamMasuk > 22 {
					fmt.Println("Jam masuk parkir tidak valid!")
					fmt.Print("Jam Masuk : ")
					fmt.Scan(&jamMasuk, &menitMasuk)
				}
				pMobil[*nMobil].waktu.jamMasuk = jamMasuk
				pMobil[*nMobil].waktu.menitMasuk = menitMasuk // jam masuk min 07.00 maks 22.59
				dataKendaraan[*nData] = pMobil[*nMobil]
				*nMobil++
				*nData++
			}
		} else if input == 2 {
			if *nMotor > MAXPARKIRMOTOR-1 {
				fmt.Println("Parkiran motor sudah penuh!")
				*nMotor = MAXPARKIRMOTOR - 1
			} else {
				pMotor[*nMotor].tipe = "Motor"
				fmt.Print("Hari (Minggu,..,Jumat,Sabtu) : ")
				fmt.Scan(&pMotor[*nMotor].hari)
				fmt.Print("Nomor Polisi : ")
				fmt.Scan(&nopol)
				for !validNoPolMobil(*pMobil, *nMobil, nopol) || !validNoPolMotor(*pMotor, *nMotor, nopol) {
					fmt.Println("Nomor Polisi tidak valid!")
					fmt.Print("Nomor Polisi :")
					fmt.Scan(&nopol)
				}
				pMotor[*nMotor].noPol = nopol
				fmt.Print("Vip : ")
				fmt.Scan(&pMotor[*nMotor].vip)
				fmt.Print("Valet : ")
				fmt.Scan(&pMotor[*nMotor].valet)
				fmt.Print("Jam Masuk : ")
				fmt.Scan(&jamMasuk, &menitMasuk)
				for jamMasuk < 7 || jamMasuk > 22 {
					fmt.Println("Jam masuk parkir tidak valid! ")
					fmt.Print("Jam Masuk : ")
					fmt.Scan(&jamMasuk, &menitMasuk)
				}
				pMotor[*nMotor].waktu.jamMasuk = jamMasuk
				pMotor[*nMotor].waktu.menitMasuk = menitMasuk
				dataKendaraan[*nData] = pMotor[*nMotor]
				*nMotor++
				*nData++
			}
		}
	}
}

func keluarParkir(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int) {
	var nopol string
	var idx, idxData, jamKeluar, menitKeluar int
	fmt.Print("Masukkan Nomor Polisi : ")
	fmt.Scan(&nopol)
	for validNoPolMobil(*pMobil, *nMobil, nopol) && validNoPolMotor(*pMotor, *nMotor, nopol) {
		fmt.Println("Nomor Polisi tidak terdapat di dalam parkiran!")
		fmt.Print("Masukkan Nomor Polisi : ")
		fmt.Scan(&nopol)
	}
	idxData = searchIdxNoPol(*dataKendaraan, *nData, nopol)
	if !validNoPolMobil(*pMobil, *nMobil, nopol) {
		idx = searchIdxNoPolMobil(*pMobil, *nMobil, nopol)
		fmt.Print("Jam Keluar : ")
		fmt.Scan(&jamKeluar, &menitKeluar)
		for (jamKeluar < dataKendaraan[idxData].waktu.jamMasuk) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar < dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar >= 23 && menitKeluar > 0) || (jamKeluar < 7) {
			fmt.Println("Jam keluar parkir tidak valid!")
			fmt.Print("Jam Keluar : ")
			fmt.Scan(&jamKeluar, &menitKeluar)
		}
		pMobil[idx].waktu.jamKeluar = jamKeluar
		pMobil[idx].waktu.menitKeluar = menitKeluar //jam keluar maks 23.00
		dataKendaraan[idxData] = pMobil[idx]
		hapusMobil(&*pMobil, &*nMobil, idx)
	} else if !validNoPolMotor(*pMotor, *nMotor, nopol) {
		idx = searchIdxNoPolMotor(*pMotor, *nMotor, nopol)
		fmt.Print("Jam Keluar : ")
		fmt.Scan(&jamKeluar, &menitKeluar)
		for (jamKeluar < dataKendaraan[idxData].waktu.jamMasuk) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar < dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar >= 23 && menitKeluar > 0) || (jamKeluar < 7) {
			fmt.Println("Jam keluar parkir tidak valid!")
			fmt.Print("Jam Keluar : ")
			fmt.Scan(&jamKeluar, &menitKeluar)
		}
		pMotor[idx].waktu.jamKeluar = jamKeluar
		pMotor[idx].waktu.menitKeluar = menitKeluar //jam keluar maks 23.00
		dataKendaraan[idxData] = pMotor[idx]
		hapusMotor(&*pMotor, &*nMotor, idx)
	}
	hitungHarga(&*dataKendaraan, idxData)
	*totalPendapatan += dataKendaraan[idxData].hargaParkir
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("Hari : %s\n", dataKendaraan[idxData].hari)
	fmt.Printf("Tipe Kendaraan : %s\n", dataKendaraan[idxData].tipe)
	fmt.Printf("Nomor Polisi : %s\n", dataKendaraan[idxData].noPol)
	fmt.Printf("Vip : %t\n", dataKendaraan[idxData].vip)
	fmt.Printf("Valet : %t\n", dataKendaraan[idxData].valet)
	fmt.Printf("Jam Masuk : %d.%d\n", dataKendaraan[idxData].waktu.jamMasuk, dataKendaraan[idxData].waktu.menitMasuk)
	fmt.Printf("Jam Keluar : %d.%d\n", dataKendaraan[idxData].waktu.jamKeluar, dataKendaraan[idxData].waktu.menitKeluar)
	fmt.Printf("Harga Parkir : Rp.%d\n", dataKendaraan[idxData].hargaParkir)
	fmt.Printf("Total Pendapatan : Rp.%d\n", *totalPendapatan)
	fmt.Println("--------------------------------------------------------------------------------------------")
}

func validNoPolMobil(pMobil tabParkirMobil, nMobil int, nopol string) bool {
	var i int
	var ketemu bool = false
	for i < nMobil && !ketemu {
		ketemu = pMobil[i].noPol == nopol
		i++
	}
	return !ketemu
}

func validNoPolMotor(pMotor tabParkirMotor, nMotor int, nopol string) bool {
	var i int
	var ketemu bool = false
	for i < nMotor && !ketemu {
		ketemu = pMotor[i].noPol == nopol
		i++
	}
	return !ketemu
}

func validNoPol(dataKendaraan tabDataKendaraan, nData int, nopol string) bool {
	var i int
	var ketemu bool = false
	for i < nData && !ketemu {
		ketemu = dataKendaraan[i].noPol == nopol
		i++
	}
	return !ketemu
}

func searchIdxNoPolMobil(pMobil tabParkirMobil, nMobil int, nopol string) int {
	var i, idx int
	idx = -1
	for i < nMobil && idx == -1 {
		if pMobil[i].noPol == nopol {
			idx = i
		}
		i++
	}
	return idx
}

func searchIdxNoPolMotor(pMotor tabParkirMotor, nMotor int, nopol string) int {
	var i, idx int
	idx = -1
	for i < nMotor && idx == -1 {
		if pMotor[i].noPol == nopol {
			idx = i
		}
		i++
	}
	return idx
}

func searchIdxNoPol(dataKendaraan tabDataKendaraan, nData int, nopol string) int {
	var i, idx int
	idx = -1
	for i < nData && idx == -1 {
		if dataKendaraan[i].noPol == nopol {
			idx = i
		}
		i++
	}
	return idx
}

func hitungHarga(dataKendaraan *tabDataKendaraan, idxData int) {
	var harga, totalMenit, totalJam int
	totalMenit = (dataKendaraan[idxData].waktu.jamKeluar*60 + dataKendaraan[idxData].waktu.menitKeluar) - (dataKendaraan[idxData].waktu.jamMasuk*60 + dataKendaraan[idxData].waktu.menitMasuk)
	totalJam = totalMenit / 60
	if totalMenit%60 != 0 {
		totalJam++
	}
	if dataKendaraan[idxData].tipe == "Mobil" {
		harga = 10000 + (totalJam-1)*5000
		if dataKendaraan[idxData].vip {
			harga += 20000
		}
		if dataKendaraan[idxData].valet {
			harga += 50000
		}
	} else if dataKendaraan[idxData].tipe == "Motor" {
		harga = 5000 + (totalJam-1)*2000
		if dataKendaraan[idxData].vip {
			harga += 10000
		}
		if dataKendaraan[idxData].valet {
			harga += 20000
		}
	}
	if dataKendaraan[idxData].hari == "Jumat" {
		harga = harga - (harga * 20 / 100)
	}
	dataKendaraan[idxData].hargaParkir = harga
}

func hapusMobil(pMobil *tabParkirMobil, nMobil *int, idx int) {
	var i int
	i = idx
	for i < *nMobil-1 {
		pMobil[i] = pMobil[i+1]
		i++
	}
	*nMobil--
}

func hapusMotor(pMotor *tabParkirMotor, nMotor *int, idx int) {
	var i int
	i = idx
	for i < *nMotor-1 {
		pMotor[i] = pMotor[i+1]
		i++
	}
	*nMotor--
}

func editKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, totalPendapatan *int, nMobil, nMotor, nData int) {
	var input int
	for input != 8 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                             *MENU EDIT DATA KENDARAAN*                             *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Data Kendaraan yang ingin diedit :")
		fmt.Println("1. Hari")
		fmt.Println("2. Tipe")
		fmt.Println("3. Nomor Polisi")
		fmt.Println("4. VIP")
		fmt.Println("5. Valet")
		fmt.Println("6. Jam Masuk")
		fmt.Println("7. Jam Keluar")
		fmt.Println("8. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5/6/7/8) : ")
		fmt.Scan(&input)
		if input == 1 || input == 2 || input == 3 || input == 4 || input == 5 || input == 6 || input == 7 {
			editDataKendaraan(&*pMobil, &*pMotor, &*dataKendaraan, &*totalPendapatan, nMobil, nMotor, nData, input)
		}
	}
}

func editDataKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, totalPendapatan *int, nMobil, nMotor, nData, input int) {
	var idxData, idxMobil, idxMotor, jamMasuk, jamKeluar, menitMasuk, menitKeluar int
	var nopol, nopolEdit string
	fmt.Print("Masukkan Nomor Polisi kendaraan yang akan diedit : ")
	fmt.Scan(&nopolEdit)
	idxData = searchIdxNoPol(*dataKendaraan, nData, nopolEdit)
	if idxData == -1 {
		fmt.Println("Nomor Polisi tersebut tidak terdapat di dalam data kendaraan!")
	} else {
		*totalPendapatan -= dataKendaraan[idxData].hargaParkir
		if input == 1 {
			fmt.Print("Hari : ")
			fmt.Scan(&dataKendaraan[idxData].hari)
			fmt.Println("Hari berhasil diedit :)")
		} else if input == 2 {
			fmt.Print("Tipe (Mobil/Motor) : ")
			fmt.Scan(&dataKendaraan[idxData].tipe)
			fmt.Println("Tipe kendaraan berhasil diedit :)")
		} else if input == 3 {
			fmt.Println("Nomor Polisi : ")
			fmt.Scan(&nopol)
			for !validNoPol(*dataKendaraan, nData, nopol) {
				fmt.Println("Nomor Polisi tidak valid!")
				fmt.Print("Nomor Polisi :")
				fmt.Scan(&nopol)
				fmt.Println("Nomor Polisi berhasil diedit :)")
			}
			dataKendaraan[idxData].noPol = nopol
		} else if input == 4 {
			fmt.Print("Vip : ")
			fmt.Scan(&dataKendaraan[idxData].vip)
			fmt.Println("Vip berhasil diedit :)")
		} else if input == 5 {
			fmt.Print("Valet : ")
			fmt.Scan(&dataKendaraan[idxData].valet)
			fmt.Println("Valet berhasil diedit :)")
		} else if input == 6 {
			fmt.Print("Jam Masuk : ")
			fmt.Scan(&jamMasuk, &menitMasuk)
			for jamMasuk < 7 || jamMasuk > 22 {
				fmt.Println("Jam masuk parkir tidak valid!")
				fmt.Print("Jam Masuk : ")
				fmt.Scan(&jamMasuk, &menitMasuk)
			}
			dataKendaraan[idxData].waktu.jamMasuk = jamMasuk
			dataKendaraan[idxData].waktu.menitMasuk = menitMasuk
			fmt.Println("Jam Masuk berhasil diedit :)")
		} else if input == 7 {
			fmt.Print("Jam Keluar : ")
			fmt.Scan(&jamKeluar, &menitKeluar)
			for (jamKeluar < dataKendaraan[idxData].waktu.jamMasuk) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar < dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar >= 23 && menitKeluar > 0) || (jamKeluar < 7) {
				fmt.Println("Jam keluar parkir tidak valid!")
				fmt.Print("Jam Keluar : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
			}
			dataKendaraan[idxData].waktu.jamKeluar = jamKeluar
			dataKendaraan[idxData].waktu.menitKeluar = menitKeluar
			fmt.Println("Jam Keluar berhasil diedit :)")
		}
		if !validNoPolMobil(*pMobil, nMobil, nopolEdit) {
			idxMobil = searchIdxNoPolMobil(*pMobil, nData, nopolEdit)
			pMobil[idxMobil] = dataKendaraan[idxData]
		} else if !validNoPolMotor(*pMotor, nMotor, nopolEdit) {
			idxMotor = searchIdxNoPolMotor(*pMotor, nData, nopolEdit)
			pMotor[idxMotor] = dataKendaraan[idxData]
		}
		hitungHarga(&*dataKendaraan, idxData)
		*totalPendapatan += dataKendaraan[idxData].hargaParkir
	}
}

func hapusDataKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int) {
	var nopol string
	var idxData, idxMobil, idxMotor, i int
	fmt.Println("============================================================================================")
	fmt.Println("*||*                             *MENU HAPUS DATA KENDARAAN*                            *||*")
	fmt.Println("============================================================================================")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan Nomor Polisi dari kendaraan yang ingin dihapus : ")
	fmt.Scan(&nopol)
	idxData = searchIdxNoPol(*dataKendaraan, *nData, nopol)
	if idxData == -1 {
		fmt.Println("Nomor Polisi tersebut tidak terdapat di dalam data kendaraan yang parkir!")
	} else {
		i = idxData
		*totalPendapatan -= dataKendaraan[i].hargaParkir
		for i < *nData-1 {
			dataKendaraan[i] = dataKendaraan[i+1]
			i++
		}
		*nData--
		if !validNoPolMobil(*pMobil, *nMobil, nopol) {
			idxMobil = searchIdxNoPolMobil(*pMobil, *nMobil, nopol)
			i = idxMobil
			for i < *nMobil-1 {
				pMobil[i] = pMobil[i+1]
				i++
			}
			*nMobil--
		} else if !validNoPolMotor(*pMotor, *nMotor, nopol) {
			idxMotor = searchIdxNoPolMotor(*pMotor, *nMotor, nopol)
			i = idxMotor
			for i < *nMotor-1 {
				pMotor[i] = pMotor[i+1]
				i++
			}
			*nMotor--
		}
		fmt.Printf("Semua data kendaraan dengan Nomor Polisi %s berhasil dihapus :)\n", nopol)
	}

}

func cariKendaraan(dataKendaraan tabDataKendaraan, nData int) {
	var nopol string
	var idxData int
	fmt.Print("Masukkan Nomor Polisi : ")
	fmt.Scan(&nopol)
	for validNoPol(dataKendaraan, nData, nopol) {
		fmt.Println("Nomor Polisi tidak terdapat di dalam data parkiran!")
		fmt.Print("Masukkan Nomor Polisi : ")
		fmt.Scan(&nopol)
	}
	idxData = searchIdxNoPol(dataKendaraan, nData, nopol)
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("Hari : %s\n", dataKendaraan[idxData].hari)
	fmt.Printf("Tipe Kendaraan : %s\n", dataKendaraan[idxData].tipe)
	fmt.Printf("Nomor Polisi : %s\n", dataKendaraan[idxData].noPol)
	fmt.Printf("Vip : %t\n", dataKendaraan[idxData].vip)
	fmt.Printf("Valet : %t\n", dataKendaraan[idxData].valet)
	fmt.Printf("Jam Masuk : %d.%d\n", dataKendaraan[idxData].waktu.jamMasuk, dataKendaraan[idxData].waktu.menitMasuk)
	fmt.Printf("Jam Keluar : %d.%d\n", dataKendaraan[idxData].waktu.jamKeluar, dataKendaraan[idxData].waktu.menitKeluar)
	fmt.Printf("Harga Parkir : Rp.%d\n", dataKendaraan[idxData].hargaParkir)
	if dataKendaraan[idxData].hargaParkir == 0 {
		fmt.Println("Status : Masih dalam parkiran")
	} else {
		fmt.Println("Status : Sudah keluar parkiran")
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
}

func urutBerdasarkanHargaDescending(dataKendaraan tabDataKendaraan, nData int) {
	var i, pass int
	var temp tKendaraan
	for pass = 1; pass <= nData-1; pass++ {
		temp = dataKendaraan[pass]
		i = pass
		for i > 0 && dataKendaraan[i-1].hargaParkir < temp.hargaParkir {
			dataKendaraan[i] = dataKendaraan[i-1]
			i--
		}
		dataKendaraan[i] = temp
	}
}

func urutBerdasarkanHargaAscending(dataKendaraan tabDataKendaraan, nData int) {
	var i, pass int
	var temp tKendaraan
	for pass = 1; pass <= nData-1; pass++ {
		temp = dataKendaraan[pass]
		i = pass
		for i > 0 && dataKendaraan[i-1].hargaParkir > temp.hargaParkir {
			dataKendaraan[i] = dataKendaraan[i-1]
			i--
		}
		dataKendaraan[i] = temp
	}
}

func urutBerdasarkanHari(dataKendaraan tabDataKendaraan, nData int, hari string) {

}

func urutberdasarkanTipe(dataKendaraan tabDataKendaraan, nData int, tipe string) {

}

func cetakKendaraan(pMobil tabParkirMobil, pMotor tabParkirMotor, dataKendaraan tabDataKendaraan, nMobil, nMotor, nData int) {
	var input int
	for input != 4 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                            *MENU CETAK DATA KENDARAAN*                             *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Data Kendaraan yang ingin dicetak :")
		fmt.Println("1. Data di Parkiran Mobil")
		fmt.Println("2. Data di Parkiran Motor")
		fmt.Println("3. Data Keseluruhan")
		fmt.Println("4. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4) : ")
		fmt.Scan(&input)
		if input == 1 {
			cetakParkirMobil(pMobil, nMobil)
		} else if input == 2 {
			cetakParkirMotor(pMotor, nMotor)
		} else if input == 3 {
			cetakDataKeseluruhan(dataKendaraan, nData)
		}
	}
}

func cetakDataKeseluruhan(dataKendaraan tabDataKendaraan, nData int) {
	var input int
	var tipe, hari string
	for input != 5 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih urutan data yang ingin dicetak :")
		fmt.Println("1. Urut berdasarkan harga tertinggi")
		fmt.Println("2. Urut berdasarkan harga terendah")
		fmt.Println("3. Urut berdasarkan tipe kendaraan")
		fmt.Println("4. Urut berdasarkan hari")
		fmt.Println("5. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5) : ")
		fmt.Scan(&input)
		if input == 1 {
			urutBerdasarkanHargaDescending(dataKendaraan, nData)
			cetakParkirKeseluruhan(dataKendaraan, nData)
		} else if input == 2 {
			urutBerdasarkanHargaAscending(dataKendaraan, nData)
			cetakParkirKeseluruhan(dataKendaraan, nData)
		} else if input == 3 {
			fmt.Println("Masukkan Tipe Kendaraan : ")
			fmt.Scan(&tipe)
			urutberdasarkanTipe(dataKendaraan, nData, tipe)
			cetakParkirKeseluruhan(dataKendaraan, nData)
		} else if input == 4 {
			urutBerdasarkanHari(dataKendaraan, nData, hari)
			cetakParkirKeseluruhan(dataKendaraan, nData)
		}
	}
}

func cetakParkirMobil(pMobil tabParkirMobil, nMobil int) {
	var i int
	fmt.Println("Data Parkiran Mobil : ")
	for i = 0; i < nMobil; i++ {
		fmt.Println(pMobil[i].tipe, pMobil[i].hari, pMobil[i].noPol, pMobil[i].vip, pMobil[i].valet, pMobil[i].waktu.jamMasuk, pMobil[i].waktu.menitMasuk, pMobil[i].waktu.jamKeluar, pMobil[i].waktu.menitKeluar)
	}
}

func cetakParkirMotor(pMotor tabParkirMotor, nMotor int) {
	var i int
	fmt.Println("Data Parkiran Motor : ")
	for i = 0; i < nMotor; i++ {
		fmt.Println(pMotor[i].tipe, pMotor[i].hari, pMotor[i].noPol, pMotor[i].vip, pMotor[i].valet, pMotor[i].waktu.jamMasuk, pMotor[i].waktu.menitMasuk, pMotor[i].waktu.jamKeluar, pMotor[i].waktu.menitKeluar)
	}
}

func cetakParkirKeseluruhan(dataKendaraan tabDataKendaraan, nData int) {
	var i int
	fmt.Println("Data Parkiran Keseluruhan")
	for i = 0; i < nData; i++ {
		fmt.Println(dataKendaraan[i].tipe, dataKendaraan[i].hari, dataKendaraan[i].noPol, dataKendaraan[i].vip, dataKendaraan[i].valet, dataKendaraan[i].waktu.jamMasuk, dataKendaraan[i].waktu.menitMasuk, dataKendaraan[i].waktu.jamKeluar, dataKendaraan[i].waktu.menitKeluar)
	}
}
