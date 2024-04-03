//buat fungsi saja
func cashback(total int, mem bool) int {
	var cb int
	if total >= 500000 && mem {
		cb = total * 20 / 100
		return cb
	} else {
		return cb
	}
}