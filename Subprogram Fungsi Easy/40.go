func cekHitung(b1 int, b2 int, b3 int, jumlah int) bool {
	//mengembalikan nilai true jika  jumlah sama dengan penjumlahan b1, b2,b3 dan false jika sebaliknya
	var cek bool
	cek = b1+b2+b3 == jumlah
	return cek

}