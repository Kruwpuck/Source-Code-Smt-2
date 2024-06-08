//isi fungsi saja
func hitungDeret(n, m int) float64 {
	var total float64
	total = 0.0
	var sign float64
	sign = 1.0
	var i int
	for i = n; i <= m; i++ {
	total = total + sign * (3.0 / float64(i))
	sign = sign * -1
	}
	return total
   }
   