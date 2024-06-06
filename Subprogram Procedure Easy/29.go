func deskripsiHuruf(huruf byte) {
	/*
	I.S. terdefinisi sebuah karakter yang menyatakan nilai huruf.
	F.S. tercetak deskripsi berdasarkan tabel.
	*/
	var nilai int
	if huruf == 'A' {
	nilai = 5
	} else if huruf == 'B' {
	nilai = 4
	} else if huruf == 'C' {
	nilai = 3
	} else if huruf == 'D' {
	nilai = 2
	} else if huruf == 'E' {
	nilai = 1
	} else if huruf == 'T' {
	nilai = 0
	}
    fmt.Println(nilai)
}