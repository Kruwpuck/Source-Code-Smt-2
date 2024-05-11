package main

import "fmt"

const LimitData int = 100

type FormatSoal struct {
	pertanyaan string
	opsi       [4]string
	jawaban    int
	correct    int
}

type TabSoal struct {
	data  [LimitData]FormatSoal
	n     int
	nShow int // untuk menampilkan berapa soal
}

type FormatUser struct {
	Username                      string
	FirstName, LastName, AsalKota string
	password                      string
	id                            int
	score                         int
}

type TabUser struct {
	data [LimitData]FormatUser
	n    int
}

func main() {
	var input_pilihan int
	var BankSoal TabSoal
	var User TabUser
	for input_pilihan != 3 {
		fmt.Println("===========================================================================================================")
		fmt.Println("===========================================================================================================")
		fmt.Println("||                                                                                                       ||")
		fmt.Println("||                                            WELCOME TO                                                 ||")
		fmt.Println("||                                      WHO ONE TO BE A MILLIONER                                        ||")
		fmt.Println("||                            Salsabila Rahayu Putri // Rifkii Aiidan Maulana                            ||")
		fmt.Println("||                                (103032330153)     //    (103032300159)                                ||")
		fmt.Println("||                                                                                                       ||")
		fmt.Println("===========================================================================================================")
		fmt.Println("===========================================================================================================")
		fmt.Println("                                                                                                           ")
		fmt.Println("1. ADMIN")
		fmt.Println("2. KONTESTAN")
		fmt.Println("3. EXIT QUIZ")
		fmt.Println("Masukkan Pilihan Anda (1/2/3 ?) :")
		for input_pilihan != 3 {
			if input_pilihan == 1 {
				menu_Admin(&BankSoal, &User)
			} else if input_pilihan == 2 {
				//1menu_kontestan(&BankSoal, &User)
			} else if input_pilihan == 3 {
				fmt.Println("")
				fmt.Println("********************************* ERROR, INPUT TIDAK DIKENALI !!!  *********************************")
				fmt.Println("")
			}
			fmt.Scan(&input_pilihan)
		}

	}
}

func menu_Admin(BankSoal *TabSoal, User *TabUser) {
	var pilihan, x string
	fmt.Println("---------------------------------------------- MENU ADMIN---------------------------------------------------")
	fmt.Println("| a. Tambah Soal                                                                                           |")
	fmt.Println("| b. Hapus Soal                                                                                            |")
	fmt.Println("| c. Edit Soal                                                                                             |")
	fmt.Println("| d. Set Jumlah Jawaban                                                                                    |")
	fmt.Println("| e. Tampilan Soal yang Paling Banyak diJawab Benar                                                        |")
	fmt.Println("| f. Tampilan Soal yang Paling Banyak diJawab Salah                                                        |")
	fmt.Println("| g. Cari kontestan berdasarkan id                                                                         |")
	fmt.Println("| h. Log Out                                                                                               |")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan Pilihan anda (a/b/c/d/e/f/g/h?):")
	fmt.Scan(&pilihan)
	if pilihan == "a" {
		tambah_soal(&*BankSoal)
	} else if pilihan == "b" {
		hapus_soal(&*BankSoal, x)
	} else if pilihan == "c" {
		edit_soal(&*BankSoal, x)
	} else if pilihan == "d" {
		set_jumlah(&*BankSoal)
	} else if pilihan == "e" {
		//Soal_Benar_Terbanyak(&*BankSoal)
	} else if pilihan == "f" {
		//Soal_Salah_Terbanyak(&*BankSoal)
	} else if pilihan == "g" {
		//cari_kontestan(&*User)
	} else if pilihan == "h" {
		fmt.Println("*")
		fmt.Println("************* ERROR, INPUT TIDAK DIKENALI !!!")
		fmt.Println("*")

	}
}

func tambah_soal(BankSoal *TabSoal) {
	var temp FormatSoal
	fmt.Println("---------------------------------------------- MENU TAMBAH SOAL---------------------------------------------------")
	fmt.Println("Masukkan Soal yang akan ditambahkan:")
	fmt.Scan(&temp.pertanyaan)
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan opsi jawaban ke-%d: ", i+1)
		fmt.Scan(&temp.opsi[i])
	}
	fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
	BankSoal.data[BankSoal.n] = temp
	BankSoal.n++
	fmt.Println("*")
	fmt.Println("SOAL BERHASIL DITAMBAHKAN KE BANK SOAL")
	fmt.Println("*")
}

func sequential_pertanyaan(BankSoal *TabSoal, x string) int {
	var i, idx int
	idx = -1
	i = 1
	for i <= BankSoal.n && idx == -1 {
		if BankSoal.data[i].pertanyaan == x {
			idx = i
		}
	}
	return idx
}

func edit_soal(BankSoal *TabSoal, x string) {
	var idx, i int
	idx = sequential_pertanyaan(BankSoal, x)
	if idx == -1 {
		fmt.Println("Tidak ada soal seperti itu")
	} else {
		fmt.Println("Masukkan Soal yang akan diedit :")
		fmt.Scan(&BankSoal.data[idx].pertanyaan)
		// untuk ganti opsi
		for i = 0; i < 4; i++ {
			fmt.Println("Masukkan Opsi Baru:")
			fmt.Scan(&BankSoal.data[idx].opsi)
		}
		fmt.Println("Masukkan Kunci Jawaban Baru [1/2/3/4]: ")
		fmt.Scan(&BankSoal.data[idx].jawaban)
	}

}

func hapus_soal(BankSoal *TabSoal, x string) {
	var idx, i int
	idx = sequential_pertanyaan(BankSoal, x)
	if idx == -1 {
		fmt.Println("Tidak Terdapat Soal Seperti itu")
	} else {
		i = idx
		for i < BankSoal.n {
			BankSoal.data[i] = BankSoal.data[i+1]
			i++
		}
		BankSoal.n--
	}
}

func set_jumlah(BankSoal *TabSoal) {
	fmt.Println("Set Jumlah Soal Quiz :")
	fmt.Scan(&BankSoal.nShow)
}
