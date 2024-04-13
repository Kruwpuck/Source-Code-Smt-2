const nMax int = 1000

type tabInt [nMax]int

func isiArray(arr1, arr2 *tabInt, n int) {
	/* I.S. Data tersedia dalam piranti masukan
	F.S. arr1, arr2 telah terisi sejumlah bilangan */
	// variabel digunakan dalam perulangan dan menyimpan sementara nilai untuk masukkan
	var i int
	for i = 0; i < n; i++ { // mengisi array 1
		fmt.Scan(&arr1[i])
		for arr1[i] != 0 && arr1[i] != 1 { // menerima inputan hingga inputan yang dimasukkan 1 atau 0
			fmt.Scan(&arr1[i])
		}

	}

	for i = 0; i < n; i++ { // mengisi array 2
		fmt.Scan(&arr2[i])
		for arr2[i] != 0 && arr2[i] != 1 { // menerima inputan hingga inputan yang dimasukkan 1 atau 0
			fmt.Scan(&arr2[i])
		}

	}
}

func and(x1, x2 int) int {
	/*mengembalikan nilai 1 apabila memenuhi logika and dan 0 jika tidak memenuhi*/
	if x1 == x2 && x1 == 1 {
		return 1
	} else {
		return 0
	}
}

func or(x1, x2 int) int {
	/*mengembalikan nilai 1 apabila memenuhi logika OR dan 0 jika tidak memenuhi*/
	if x2 == 1 || x1 == 1 {
		return 1
	} else {
		return 0
	}
}

func hitungOR(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi or yang memenuhi kondisi pada array*/
	var total int // variabel digunakan untuk iterasi loop dan menyimpan total banyaknya operasi yang memenuhi
	for i := 0; i < n; i++ {
		total += or(arr1[i], arr2[i])
	}
	return total
}

func hitungAND(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi logika and yang memenuhi kondisi pada array*/
	var total int
	for i := 0; i < n; i++ {
		total += and(arr1[i], arr2[i])
	}
	return total
}

func xor(x1, x2 int) int {
	/*mengembalikan nilai 1 apabila memenuhi logika XOR dan 0 jika tidak memenuhi*/
	if x1 != x2 { // syarat ketika memenuhi XOR yaitu ketika nilai x1 != nilai x2
		return 1
	} else {
		return 0
	}
}

func hitungXOR(arr1, arr2 tabInt, n int) int {
	/*mengembalikan total dari banyaknya operasi xor yang memenuhi kondisi pada array*/
	var total int
	for i := 0; i < n; i++ {
		total += xor(arr1[i], arr2[i])
	}
	return total
}

func cetak(arr1, arr2 tabInt, n int) {
	/* I.S. array arr1,arr2 berisi sejumlah n bilangan 1 dan 0
	F.S. Banyaknya hasil dari operasi logika OR, hasil dari operasi logika OR, banyaknya hasil dari operasi logika AND
	hasil dari operasi logika AND, banyaknya hasil dari operasi logika XOR,hasil dari operasi logika OR ditampilkan di layar */
	var i int
	fmt.Println("Total OR:", hitungOR(arr1, arr2, n)) // menampilkan total OR yang memenuhi
	for i = 0; i < n; i++ {
		fmt.Print(or(arr1[i], arr2[i]), " ")
	}
	fmt.Println("")
	fmt.Println("Total AND:", hitungAND(arr1, arr2, n))
	for i = 0; i < n; i++ {
		fmt.Print(and(arr1[i], arr2[i]), " ")
	}
	fmt.Println("")
	fmt.Println("Total XOR:", hitungXOR(arr1, arr2, n)) // menampilkan total operasi logika XOR yang memenuhi
	for i = 0; i < n; i++ {
		fmt.Print(xor(arr1[i], arr2[i]), " ")
	}
	fmt.Println()
}
