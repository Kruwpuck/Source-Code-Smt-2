func UnGeometri(a, r, n int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif a, r dan n sebagai jumlah.
	 F.S. Fungsi mengembalikan hasil nilai dari Suku ke-n dari deret geometri */
	 if n == 1 {
	 return a
	 } else {
	 return UnGeometri(a, r, n - 1) * r// gunakan baris ini untuk memproses per
	 }
	}