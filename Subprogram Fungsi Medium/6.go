// implementasikan fungsinya saja
func jumlahDeret(x, y int) float64 {
	/* mengembalikan nilai dari jumlah deret n dan m */
	 var i int
	 var hasil float64
	 for i = x; i <= y; i++{
	 hasil = hasil + 2/float64(i)
	 }
	 return hasil
	}