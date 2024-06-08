func sumCubic(n int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	 F.S. Fungsi mengembalikan hasil perpangkatan 3 */
	 if n == 1 {
	 return 1
	 } else {
	 return (n * n *n) + sumCubic(n - 1)// gunakan baris ini untuk memproses
	 }
	
	}