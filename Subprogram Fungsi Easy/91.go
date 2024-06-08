// Lengkapi fungsi di bawah

func pemenang(x3, x2, x1, y3, y2, y1 int) string {
	// Mengembalikan string "A", "B", atau "imbang" berdasarkan
	// selisih total nilai antara tim A dengan tim B
	var jmlX, jmlY int
	jmlX = x3*3 + x2*2 + x1
	jmlY = y3*3 + y2*2 + y1
	if jmlX > jmlY {
		return "A"
	} else if jmlY > jmlX {
		return "B"
	} else {
		return "imbang"
	}
}