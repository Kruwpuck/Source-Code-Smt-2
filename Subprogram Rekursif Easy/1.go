func jumlah(n int)int{
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	 F.S. Fungsi mengembalikan hasil penjumlahan 1, 2, 3, ...., n. }*/
	 if n == 1{
	 return 1
	 }else {
	 return jumlah(n-1) + n
	 }
	}
	