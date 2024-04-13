const NMax int = 1000

type tabInt [NMax]int

func masukanArray(T *tabInt, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat n. Data string sebanyak n buah tersedia
	  pada input device*/
	/*F.S. Array T berisi data yang diberikan*/
	/*Petunjuk : Lakukan perulangan sebanyak n kali untuk melakukan input terhadap
	  data nilai mahasiswa*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func isDescending(T tabInt, n int) bool {
	/*Mengembalikan boolean true jika isi data pada array terurut secara
	  descending, boolean false jika tidak*/
	descen := true
	for i := 0; i < n-1; i++ {
		if T[i] > T[i+1] {
			descen = true
		} else {
			descen = false
			break
		}
	}
	return descen
}

func isAscending(T tabInt, n int) bool {
	/*Mengembalikan boolean true jika isi data pada array terurut secara
	  ascending, boolean false jika tidak*/
	ascen := true
	for i := 0; i < n-1; i++ {
		if T[i] < T[i+1] {
			ascen = true
		} else {
			ascen = false
			break
		}
	}
	return ascen
}
func cetakKeterurutan(T tabInt, n int) {
	/*I.S. Terdefinisi array T yang berisikan sejumlah n bilangan bulat
	  F.S. Menampilkan pada layar string "Data pada array terurut secara descending"
	  atau "Data pada array terurut secara ascending" atau "Data pada array tidak
	  terurut sama sekali" menyesuaikan dengan kondisi pada array.
	  Petunjuk: Manfaatkan function isAscending dan isDescending
	*/
	if isAscending(T, n) {
		fmt.Println("Data pada array terurut secara ascending")
	} else if isDescending(T, n) {
		fmt.Println("Data pada array terurut secara descending")
	} else {
		fmt.Println("Data pada array tidak terurut sama sekali")
	}
}