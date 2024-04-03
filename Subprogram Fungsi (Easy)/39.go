func sumRepeatUntil(x int) int {
	//Mengembalikan hasil penjumlahan sesuai dengan soal (gunakan repeat untill)
	var stop bool
	var sum, i int
	stop = false
	sum = 0
	i = 1
	for !stop {
		sum += i
		stop = i == x
		i += 1
	}
	return sum
}