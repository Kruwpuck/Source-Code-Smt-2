func hitungLowercase(kata string, index int) int {
	/*{ I.S. Terdefinisi parameter kata sebagai input string
	F.S. menampilkan jummlah jumlah huruf lowercase dari kata menggunakan fungsi */
	if index == len(kata) { // masukan kondisi untuk mengecek index dengan jum
	return 0
	}
	if 'a' <= kata[index] && kata[index] <= 'z' {
	return 1 + hitungLowercase(kata, index+1) // gunakan fungsi rekursif yang
	}
	return hitungLowercase(kata, index+1)
   }