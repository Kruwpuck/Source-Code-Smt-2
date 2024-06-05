// Buatlah fungsi saja

func persamaanGaris(m, c float64) string {
	/*  Fungsi mengembalikan string berupa "melewati" apabila nilai c sama dengan nol atau
	"tidak melewati" apabila nilai c tidak sama dengan nol */
	if m*c == 0 {
		return "melewati"
	} else {
		return "tidak melewati"
	}

}