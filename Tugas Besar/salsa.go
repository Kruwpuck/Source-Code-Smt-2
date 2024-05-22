package main

import (
	"fmt"
	"math/rand"
	"time"
)

const LimitData int = 100

type FormatSoal struct {
	no         int
	pertanyaan string
	opsi       [4]string
	jawaban    int
	correct    int
	wrong      int
}

type TabSoal struct {
	n    int
	data [LimitData]FormatSoal
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
	use_temp(&BankSoal, &User)

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
		fmt.Print("Masukkan Pilihan Anda (1/2/3 ?) :")
		fmt.Scan(&input_pilihan)
		if input_pilihan == 1 {
			menu_Admin(&BankSoal, &User)
		} else if input_pilihan == 2 {
			menu_kontestan(&BankSoal, &User)
		} else if input_pilihan == 3 {
			fmt.Println("")
			fmt.Println("********************************* ERROR, INPUT TIDAK DIKENALI !!!  *********************************")
			fmt.Println("")
		}

	}

}

func menu_Admin(BankSoal *TabSoal, User *TabUser) {
	var pilihan string
	for pilihan != "i" {
		fmt.Println("---------------------------------------------- MENU ADMIN---------------------------------------------------")
		fmt.Println("| a. Tambah Soal                                                                                           |")
		fmt.Println("| b. Hapus Soal                                                                                            |")
		fmt.Println("| c. Edit Soal                                                                                             |")
		fmt.Println("| d. Lihat Soal                                                                                            |")
		fmt.Println("| e. Tampilan Soal yang Paling Banyak diJawab Benar                                                        |")
		fmt.Println("| f. Tampilan Soal yang Paling Banyak diJawab Salah                                                        |")
		fmt.Println("| g. Cari kontestan berdasarkan id                                                                         |")
		fmt.Println("| h. Cetak                                                                                                 |")
		fmt.Println("| i. log out                                                                                                |")
		fmt.Println("-------------------------------------------------------------------------------------- ----------------------")
		fmt.Print("Masukkan Pilihan anda (a/b/c/d/e/f/g/h/i?):")
		fmt.Scan(&pilihan)
		if pilihan == "a" {
			tambah_soal(&*BankSoal)
		} else if pilihan == "b" {
			hapus_soal(&*BankSoal)
		} else if pilihan == "c" {
			main_edit_soal(&*BankSoal)
		} else if pilihan == "d" {
			//lihat_soal(&*BankSoal)
		} else if pilihan == "e" {
			Soal_Benar_Terbanyak(&*BankSoal)
		} else if pilihan == "f" {
			Soal_Salah_Terbanyak(&*BankSoal)
		} else if pilihan == "g" {
			cari_kontestan(*User)
		} else if pilihan == "h" {
			cetak(*BankSoal)
		} else if pilihan == "i" {
			fmt.Println("*")
			fmt.Println("************* ERROR, INPUT TIDAK DIKENALI !!!")
			fmt.Println("*")

		}
	}
}

func tambah_soal(BankSoal *TabSoal) {

	var temp FormatSoal
	fmt.Println("--------------------------------------- MENU TAMBAH SOAL----------------------------------------------")
	fmt.Println("Masukkan Soal yang akan ditambahkan:")
	fmt.Scan(&temp.pertanyaan)
	temp.no = BankSoal.n + 1
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan opsi jawaban ke-%d: ", i+1)
		fmt.Scan(&temp.opsi[i])
	}
	fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
	fmt.Scan(&temp.jawaban)
	BankSoal.data[BankSoal.n] = temp
	BankSoal.n++
	fmt.Println("*")
	fmt.Println("SOAL BERHASIL DITAMBAHKAN KE BANK SOAL")
	fmt.Println("*")
}

func cetak(BankSoal TabSoal) {
	for i := 0; i < BankSoal.n; i++ {
		fmt.Println(BankSoal.data[i].no)
		fmt.Println(BankSoal.data[i].pertanyaan)
		for j := 0; j < 4; j++ {
			fmt.Println(BankSoal.data[i].opsi[j])
		}
		fmt.Println(BankSoal.data[i].jawaban)
		fmt.Println()
	}
}

func sequential_pertanyaan(BankSoal *TabSoal, no int) int {
	var i, idx int
	idx = -1
	i = 0
	for i <= BankSoal.n && idx == -1 {
		if BankSoal.data[i].no == no {
			idx = i
		}
		i++
	}
	return idx
}

func main_edit_soal(BankSoal *TabSoal) {
	// CARI SOAL
	var no, idx int
	var temp FormatSoal
	fmt.Println("Masukkan No pertanyaan yang akan diedit :")
	fmt.Scan(&no)
	idx = sequential_pertanyaan(BankSoal, no)
	if idx == -1 {
		fmt.Println("Tidak terdapat soal seperti itu")
	} else {
		fmt.Println("Masukkan Pertanyaan Baru :")
		fmt.Scan(&temp.pertanyaan)
		temp.no = BankSoal.data[idx].no
		for i := 0; i < 4; i++ {
			fmt.Printf("Masukkan opsi jawaban ke-%d: ", i+1)
			fmt.Scan(&temp.opsi[i])
		}
		fmt.Println("Masukkan Kunci Jawab [1/2/3/4]: ")
		fmt.Scan(&temp.jawaban)
		BankSoal.data[idx] = temp
		fmt.Println("*")
		fmt.Println("SOAL BERHASIL DITAMBAHKAN KE BANK SOAL")
		fmt.Println("*")
	}

}

func hapus_soal(BankSoal *TabSoal) {
	var idx, i int
	var no int
	fmt.Println("Masukkan No soal :")
	fmt.Scan(&no)
	idx = sequential_pertanyaan(BankSoal, no)
	if idx == -1 {
		fmt.Println("Tidak Terdapat Soal Seperti itu")
	} else {
		i = idx
		for i < BankSoal.n-1 {
			BankSoal.data[i] = BankSoal.data[i+1]
			i++
		}
		BankSoal.n--
	}

}

func Soal_Salah_Terbanyak(BankSoal *TabSoal) {
	var maks, maksNo int
	var i int
	maks = 0
	for i = 0; i < BankSoal.n; i++ {
		if BankSoal.data[i].wrong > maks {
			maks = BankSoal.data[i].wrong
			maksNo = BankSoal.data[i].no
		}
	}
	fmt.Println(maksNo)

}

func menu_kontestan(BankSoal *TabSoal, User *TabUser) {
	var input_user int

	for input_user != 3 {
		fmt.Println("-------------------------------------------------------------------------------------------------")
		fmt.Println("                                       MENU KONTESTAN                                            ")
		fmt.Println("-------------------------------------------------------------------------------------------------")
		fmt.Println(" 1. LOG IN ")
		fmt.Println(" 2. SIGN IN")
		fmt.Println(" 3. LOG OUT")
		fmt.Println(" Masukkan Pilihan Anda (1/2/3?):")
		fmt.Scan(&input_user)
		if input_user == 1 {
			// panggil fungsi log in
			log_in(&*BankSoal, &*User)
		} else if input_user == 2 {
			// penggil fungsi sign in
			sign_in(&*User)
		} else {
			fmt.Println("input tidak dikenali")
		}
	}
	fmt.Println("-------------------------------------------------------------------------------------------------")
}

func log_in(BankSoal *TabSoal, User *TabUser) {
	var temp FormatUser
	var status bool = false
	var idx int
	if BankSoal.n == 0 {
		fmt.Println("Belum ada kontestan yang terdaftar")
		return
	}
	for !status {
		fmt.Print("Username : ")
		fmt.Scan(&temp.Username)
		fmt.Print("Password : ")
		fmt.Scan(&temp.password)
		if ! cekPassword(*User, temp.Username, temp.password, &idx) {
			fmt.Print("Password salah, ulangi\n")
		} else {
			fmt.Print("Log in berhasil, menuju ke menu utama\n")
			menu_home_kontestan(&*BankSoal, &*User, idx)
			status = true
		}
	}
}

func sign_in(User *TabUser) {
	var temp FormatUser
	var validasi_Username bool
	var x string
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("|                                           MENU SIGN IN                                        |")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("Masukkan UserName: ")
	fmt.Scan(&x)
	validasi_Username = cekUserName(User, x)
	for validasi_Username {
		fmt.Println("username sudah digunakan, silakan masukan username yang baru")
		fmt.Scan(&x)
	}
	temp.Username = x

	fmt.Println("Masukkan FirstName: ")
	fmt.Scan(&temp.FirstName)
	fmt.Println("Masukkan LastName: ")
	fmt.Scan(&temp.LastName)
	fmt.Println("Masukkan Asal Kota: ")
	fmt.Scan(&temp.AsalKota)
	fmt.Println("Masukkan Password: ")
	fmt.Scan(&temp.password)
	fmt.Println(" Akun Anda Berhasil diBuat")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	
	User.n++
	temp.id = User.n
	User.data[User.n-1] = temp

}

func cekPassword(User TabUser, usn, pass string, idx *int) bool {
	var i int
	for i = 0; i < User.n; i++ {
		if User.data[i].Username == usn {
			if User.data[i].password == pass {
				*idx = User.data[i].id
				return true
			}

		}
	}
	return false
}

func cekUserName(User *TabUser, x string) bool {
	var i int
	var ketemu bool
	ketemu = false
	i = 0
	for i < User.n && !ketemu {
		ketemu = User.data[i].Username == x
		i++
	}
	return ketemu

}

func menu_home_kontestan(BankSoal *TabSoal, User *TabUser, idx int) {
	var input int
	for input != 3 {
		fmt.Println("*************************************************************************************************")
		fmt.Println("                                             HOME KONSTESTAN                                     ")
		fmt.Println("*************************************************************************************************")
		fmt.Println("1. Mulai Quiz")
		fmt.Println("2. Lihat Skor")
		fmt.Println("3. Log out")
		fmt.Scan(&input)
		if input == 1 {
			//fungsi mulai quiz
			User.data[idx-1].score = 0
			shuffleQuestions(BankSoal)
			kerjakan_soal(BankSoal, User, idx)
		} else if input == 2 {
			fmt.Println("Nilai anda adalah:", User.data[idx-1].score)

		} else {
			fmt.Println("Anda telah keluar dari Quiz yang dimainkan")
			break
		}

	}

}

func shuffleQuestions(BankSoal *TabSoal) {
	rand.Seed(time.Now().UnixNano())
	for i := range BankSoal.data[:BankSoal.n] {
		j := rand.Intn(i + 1)
		BankSoal.data[i], BankSoal.data[j] = BankSoal.data[j], BankSoal.data[i]
	}
}

func kerjakan_soal(BankSoal *TabSoal, User *TabUser, idx int) {
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("                                    SELAMAT MENGERJAKAN:))                                       ")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	var i, j int
	var answer FormatSoal
	if BankSoal.n > LimitData {
		BankSoal.n = LimitData
	}
	for i = 0; i < BankSoal.n; i++ {
		fmt.Printf("%d. %s\n", i+1, BankSoal.data[i].pertanyaan)
		for j = 0; j < 4; j++ {
			fmt.Printf("Opsi %d - %s\n", j+1, BankSoal.data[i].opsi[j])
		}
		fmt.Println("Masukan Jawaban (1/2/3/4?):")
		fmt.Scan(&answer.jawaban)
		if answer.jawaban == BankSoal.data[i].jawaban {
			fmt.Println("----------------------------------------------------------------------------------------")
			fmt.Println("                          Selamat Jawaban Anda Benar                                    ")
			fmt.Println("----------------------------------------------------------------------------------------")
			BankSoal.data[i].correct++
			User.data[idx-1].score += 10
		} else if answer.jawaban != BankSoal.data[i].jawaban {
			fmt.Println("----------------------------------------------------------------------------------------")
			fmt.Println("                             Jawaban Anda Salah                                         ")
			fmt.Printf("Jawaban yang Benar adalah : Opsi %d", BankSoal.data[i].jawaban)
			fmt.Println("----------------------------------------------------------------------------------------")
			BankSoal.data[i].wrong++
		}

	}
}

func lihat_skor(BankSoal *TabSoal, User *TabUser, x int) {
	fmt.Println("************************************************************************************************")
	fmt.Println("*************************                 DAFTAR SKOR             ******************************")
	fmt.Println("************************************************************************************************")
	// x = cekSkor(*BankSoal, *User)
	for i := 0; i < User.n; i++ {
		fmt.Printf("Username: %s\n", User.data[i].Username)
		fmt.Printf("Skor: %d\n", User.data[i].score)
		fmt.Println("----------------------------------------------")
	}
}

func cek_id(User TabUser, id int) int {
	var i, idx int
	idx = -1
	for i < User.n && idx == -1 {
		if User.data[i].id == id {
			idx = i
		}
	}
	return idx

}

func cari_kontestan(User TabUser) {
	var x int
	var id int
	fmt.Print("Masukkan Id : ")
	fmt.Scan(&id)
	x = cek_id(User, id)
	if x != -1 {
		fmt.Println(User.data[x].Username, User.data[x].password, User.data[x].FirstName, User.data[x].LastName)
	}

}

func Soal_Benar_Terbanyak(BankSoal *TabSoal) {
	var i, idx, temp, pass int
	for i = 1; i < BankSoal.n; i++ {
		idx = pass - 1
		for i = pass; i < BankSoal.n; i++ {
			if BankSoal.data[i].correct > BankSoal.data[idx].correct {
				idx = i
			}
		}
	}
	temp = BankSoal.data[idx].correct
	BankSoal.data[idx].correct = BankSoal.data[pass-1].correct
	BankSoal.data[pass-1].correct = temp

	for i = 0; i < BankSoal.n; i++{
		fmt.Println(BankSoal.data[temp].no, BankSoal.data[temp].pertanyaan)
	}
	
}



func use_temp(BankSoal *TabSoal, User *TabUser) {
	BankSoal.data[0].no = 1
	BankSoal.data[0].pertanyaan = "Siapakah penemu bola lampu? "
	BankSoal.data[0].opsi[0] = "Thomas Edison "
	BankSoal.data[0].opsi[1] = "Isaac Newton "
	BankSoal.data[0].opsi[2] = "Albert Einstein"
	BankSoal.data[0].opsi[3] = "Nikola Tesla "
	BankSoal.data[0].jawaban = 1

	BankSoal.data[1].no = 2
	BankSoal.data[1].pertanyaan = "Berapakah jumlah planet dalam tata surya kita? "
	BankSoal.data[1].opsi[0] = "10  "
	BankSoal.data[1].opsi[1] = "5"
	BankSoal.data[1].opsi[2] = "7"
	BankSoal.data[1].opsi[3] = "8"
	BankSoal.data[1].jawaban = 4

	BankSoal.data[2].no = 3
	BankSoal.data[2].pertanyaan = "Apa nama tokoh yang terkenal dengan eksperimen mengenai gravitasi dengan apel? "
	BankSoal.data[2].opsi[0] = "Charles Darwin "
	BankSoal.data[2].opsi[1] = "Isaac Newton"
	BankSoal.data[2].opsi[2] = "Galileo Galilei"
	BankSoal.data[2].opsi[3] = "Albert Einstein "
	BankSoal.data[2].jawaban = 2

	BankSoal.data[3].no = 4
	BankSoal.data[3].pertanyaan = "Apa nama sungai terpanjang di dunia? "
	BankSoal.data[3].opsi[0] = "Nil"
	BankSoal.data[3].opsi[1] = "Amazon"
	BankSoal.data[3].opsi[2] = "Mississippi"
	BankSoal.data[3].opsi[3] = "Yangtze"
	BankSoal.data[3].jawaban = 1

	BankSoal.data[4].no = 5
	BankSoal.data[4].pertanyaan = "Apa nama benua terbesar di dunia? "
	BankSoal.data[4].opsi[0] = "Amerika "
	BankSoal.data[4].opsi[1] = "Eropa "
	BankSoal.data[4].opsi[2] = "Asia"
	BankSoal.data[4].opsi[3] = "Afrika"
	BankSoal.data[4].jawaban = 3

	BankSoal.data[5].no = 6
	BankSoal.data[5].pertanyaan = "Apa yang merupakan simbol kimia untuk air?"
	BankSoal.data[5].opsi[0] = "O2"
	BankSoal.data[5].opsi[1] = "CO2 "
	BankSoal.data[5].opsi[2] = "H20 "
	BankSoal.data[5].opsi[3] = "H2SO4"
	BankSoal.data[5].jawaban = 3

	BankSoal.data[6].no = 7
	BankSoal.data[6].pertanyaan = "Siapakah pendiri perusahaan Microsoft? "
	BankSoal.data[6].opsi[0] = "Steve Jobs"
	BankSoal.data[6].opsi[1] = "Bill Gates"
	BankSoal.data[6].opsi[2] = "Jeff Bezos"
	BankSoal.data[6].opsi[3] = "Larry Page"
	BankSoal.data[6].jawaban = 2

	BankSoal.data[7].no = 8
	BankSoal.data[7].pertanyaan = "Apa yang dimaksud dengan AI dalam konteks Teknologi Informasi? "
	BankSoal.data[7].opsi[0] = "Advanced Interface "
	BankSoal.data[7].opsi[1] = "Automated Intelligence"
	BankSoal.data[7].opsi[2] = "Artificial Intelligence"
	BankSoal.data[7].opsi[3] = "Adaptive Innovation "
	BankSoal.data[7].jawaban = 3

	BankSoal.data[8].no = 9
	BankSoal.data[8].pertanyaan = "Apa yang dimaksud dengan Cloud Computing ?"
	BankSoal.data[8].opsi[0] = "Pengolahan data menggunakan komputer berbentuk awan"
	BankSoal.data[8].opsi[1] = "Penyimpanan data secara lokal"
	BankSoal.data[8].opsi[2] = "Pengolahan data menggunakan komputer kuantum"
	BankSoal.data[8].opsi[3] = "Pengolahan data menggunakan server terpisah via internet "
	BankSoal.data[8].jawaban = 4

	BankSoal.data[9].no = 10
	BankSoal.data[9].pertanyaan = "Apa perbedaan antara RAM dan ROM dalam komputer? "
	BankSoal.data[9].opsi[0] = "RAM bersifat permanen sedangkan ROM bersifat sementara "
	BankSoal.data[9].opsi[1] = "RAM bersifat sementara sedangkan ROM bersifat permanen"
	BankSoal.data[9].opsi[2] = "RAM dan ROM sama-sama bersifat permanen"
	BankSoal.data[9].opsi[3] = "RAM dan ROM sama-sama bersifat sementara"
	BankSoal.data[9].jawaban = 2

	BankSoal.n = 10

	User.data[0].Username = "salsa77"
	User.data[0].password = "sal07241"
	User.data[0].FirstName = "Putri"
	User.data[0].LastName = "rahayu"
	User.data[0].AsalKota = "Bandung"
	User.data[0].id = 1

	User.data[1].Username = "aisyah04"
	User.data[1].password = "88888888"
	User.data[1].FirstName = "aisyah"
	User.data[1].LastName = "putri"
	User.data[1].AsalKota = "yogyakarta"
	User.data[1].id = 2

	User.data[2].Username = "hanif08"
	User.data[2].password = "87654321"
	User.data[2].FirstName = "hanif"
	User.data[2].LastName = "fauzi"
	User.data[2].AsalKota = "palembang"
	User.data[2].id = 3

	User.data[3].Username = "jokowi01"
	User.data[3].password = "widodo01"
	User.data[3].FirstName = "jokowi"
	User.data[3].LastName = "dodo"
	User.data[3].AsalKota = "solo"
	User.data[3].id = 4

	User.data[4].Username = "megawati02"
	User.data[4].password = "karno567"
	User.data[4].FirstName = "megawati"
	User.data[4].LastName = "soekarno"
	User.data[4].AsalKota = "bali"
	User.data[4].id = 5

	User.data[5].Username = "lina09"
	User.data[5].password = "papaya50"
	User.data[5].FirstName = "lina"
	User.data[5].LastName = "kurnia"
	User.data[5].AsalKota = "bandar lampung"
	User.data[5].id = 6

	User.data[6].Username = "fatimah07"
	User.data[6].password = "sicantik2"
	User.data[6].FirstName = "fatimah"
	User.data[6].LastName = "nasution"
	User.data[6].AsalKota = "semarang"
	User.data[6].id = 7

	User.data[7].Username = "msmith02"
	User.data[7].password = "ganteng1"
	User.data[7].FirstName = "mary"
	User.data[7].LastName = "smith"
	User.data[7].AsalKota = "surabaya"
	User.data[7].id = 8

	User.data[8].Username = "bwijaya03"
	User.data[8].password = "bud09876"
	User.data[8].FirstName = "budi"
	User.data[8].LastName = "wijaya"
	User.data[8].AsalKota = "medan"
	User.data[8].id = 9

	User.data[9].Username = "tariq05"
	User.data[9].password = "abcd1234"
	User.data[9].FirstName = "tariq"
	User.data[9].LastName = "ahmad"
	User.data[9].AsalKota = "makassar"
	User.data[9].id = 10

	User.n = 10

}
