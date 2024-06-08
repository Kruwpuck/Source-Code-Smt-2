func gunung(x, y, n int) int {
	//fungsi akan mengembalikan jumlah letusan
	//return merupakan berupa bilangan bulat yang menyatakan jumlah letusan.
	var i int
	var stop = false
   
	for !stop{
	n = n - x
	i++
	stop = n - y < 0
	for !stop{
	n = n - y
	i++
	stop = n - x < 0
	break
	}
	}
	return i
   }
   