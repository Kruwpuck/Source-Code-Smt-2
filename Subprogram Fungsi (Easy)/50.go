// Buatlah fungsi saja

func manager(m1, m2, m3, m4, m5 string) string {
	/*  Fungsi mengembalikan string berupa "dipecat" atau "tidak dipecat" sesuai dengan ketentuan soal */
	if m1 == "kalah" && m2 == "kalah" && m3 == "kalah" && m4 == "kalah" && m5 == "kalah" {
		return "dipecat"
	} else {
		return "tidak dipecat"
	}
}