func kali(n int)int{
	/*{Terdefinisi sebuah bilangan bulat positif n.
	Fungsi mengembalikan hasil penjumlahan 1, 2, 3, ...., n. }*/
	if n == 1{
	return n
	}else{
	return kali(n-1) * n
	}
   }