//buat fungsi saja
func hitungD(a, b, c float64) float64{
	var diskriminan float64
	diskriminan = (b*b) - (4*a*c)
	return diskriminan
   }
   func cekTipot(hasil float64) string{
	if hasil == 0 || hasil > 0{
	return "memiliki titik potong dengan sumbu-x"
	}else {
	return "tidak memiliki titik potong dengan sumbu-x"
	}
   }