// Buat fungsi saja
func segitigaSamaSisi(a, b, c float64) bool {
	// Mengembalikan nilai boolean jika 3 bilangan real dapat
	// membentuk segitiga sama sisi
	var segitiga bool
	if a == b && a == c && a != 0 && b != 0 && c != 0 {
	segitiga = true
	}else {
	segitiga = false
	}
	return segitiga
   }
   