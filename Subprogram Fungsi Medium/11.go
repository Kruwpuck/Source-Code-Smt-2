// implementasikan fungsinya saja
func pertemuan(a, b int) int {
	/* mengembalikan jumlah pertemuan sesuai ketentuan soal */
	 total := 0
	 if a > 0 && b > 0 && a < b {
	 for i:=1;i<=365;i++ {
	 if i % a == 0 && i % b != 0 {
	 total++
	 }
	 }
	 }
	 return total
	}
	