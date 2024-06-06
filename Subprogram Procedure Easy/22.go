// buatlah prosedurnya saja
func namaHari(bilangan int) {
	/*
	I.S. terdefinisi sebuah bilangan bulat bilangan.
	F.S. tercetak nama hari yang sesuai dengan bilangan, dari 1 (Minggu) hingg
	*/
	var hari string
	if bilangan == 1 {
	hari = "Minggu"
	} else if bilangan == 2 {
	hari = "Senin"
	} else if bilangan == 3 {
	hari = "Selasa"
	} else if bilangan == 4 {
	hari = "Rabu"
	} else if bilangan == 5 {
	hari = "Kamis"
	} else if bilangan == 6 {
	hari = "Jumat"
	} else if bilangan == 7 {
	hari = "Sabtu"
	}
	fmt.Println(hari)
}