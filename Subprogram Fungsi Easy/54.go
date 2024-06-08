//buat fungsi saja
func dendaPerpustakaan(late int) int {
	var denda int
	if late > 0 && late <= 10 {
		denda = late * 2500
	} else if late > 10 {
		denda = late * 5000
	}
	return denda

}