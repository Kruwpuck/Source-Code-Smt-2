// buat prosedurnya saja

func hitungSewaBus(k, p int) {
	/*
	I.S menerima dua bilangan bulat k (kapasitas bus) dan p (jumlah peserta).
	F.S menghitung dan mencetak jumlah bus yang diperlukan untuk mengangkut 
	semua peserta dengan membagi jumlah peserta (p) 
	dengan kapasitas bus (k), dan menambahkan satu jika ada sisa peserta yang 
	tidak dapat diakomodasi penuh oleh bus.
	*/
	var jumbus int
	jumbus = p/k
	if p % k != 0{
	    jumbus += 1
	}
	fmt.Print(jumbus)
}