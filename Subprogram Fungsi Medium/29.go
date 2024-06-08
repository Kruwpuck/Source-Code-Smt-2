func belanja(no, banyak int ) int {
	/* Fungsi mengembalikan nilai dari total harga yang harus dibayar */
	var total int
	if no == 1 {
	total = banyak * 2980
	}else if no == 2{
	total = banyak * 4500
	}else if no == 3{
	total = banyak * 9980
	}else if no == 4{
	total = banyak * 4490
	}else if no == 5{
	total = banyak * 6870
	}
	if banyak < 0 {
	total = 0
	}
	return total
   }