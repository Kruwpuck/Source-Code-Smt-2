// Buatlah fungsi saja

func skor(lawan, sendiri int) string {
	/*  Fungsi mengembalikan string berupa "menang" atau "kalah" atau "draw" sesuai dengan ketentuan soal */
	if lawan > sendiri {
		return "menang"
	} else if lawan < sendiri {
		return "kalah"
	} else {
		return "draw"
	}
}