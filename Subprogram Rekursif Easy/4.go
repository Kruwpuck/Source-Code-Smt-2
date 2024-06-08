func UnAritmatika(a, b, n int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	 F.S. Fungsi mengembalikan hasil aritmatika */
	 if n == 1 {
	 return a
	 } else {
	 return UnAritmatika(a, b, n-1) + b // gunakan baris ini untuk memproses pe
	 }
	}